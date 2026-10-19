package main

import (
	"bytes"
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/yuin/goldmark/v2/ast"
	"github.com/yuin/goldmark/v2/parser"
	"github.com/yuin/goldmark/v2/text"
)

var (
	errNoEndMarker            = fmt.Errorf("no matching end marker found for start marker")
	versionSubstitutionRegexp = regexp.MustCompile(`<.+VERSION>`)
	kvPairRegexp              = regexp.MustCompile(`([a-zA-Z_]+)=(?:"(.*?[^\\])"|([^ ].+?))`)
	admonitionOpenRegexp      = regexp.MustCompile(`{{< admonition type="[^"]+?" >}}`)
	admonitionCloseRegexp     = regexp.MustCompile(`{{< /admonition >}}`)
)

func isMarker(node ast.Node, source []byte, marker string) bool {
	switch node := node.(type) {
	case *ast.Text:
		if strings.TrimSpace(node.Value.Str(source)) == marker {
			return true
		}
	case *ast.HTMLBlock, *ast.Paragraph:
		raw := rawText(node, source)
		if string(bytes.TrimSpace(raw)) == marker {
			return true
		}
	}

	return false
}

func rawText(node ast.Node, source []byte) []byte {
	buf := &bytes.Buffer{}

	switch n := node.(type) {
	case *ast.Paragraph:
		for _, line := range n.Source() {
			buf.Write(line.Bytes(source))
		}
	case *ast.HTMLBlock:
		buf.Write(n.Value.Bytes(source))
	}

	return buf.Bytes()
}

type ActionTransformer struct {
	Kind string
}

// Transform implements the parser.ASTTransformer interface and adds action metadata to any fenced code blocks within between the start and end markers.
func (t *ActionTransformer) Transform(node *ast.Document, reader text.Reader, _ parser.Context) {
	source := reader.Source()

	var (
		startMarker string
		endMarker   string
	)

	switch t.Kind {
	case "copy":
		startMarker = copyStartMarker
		endMarker = copyEndMarker
	case "exec":
		startMarker = execStartMarker
		endMarker = execEndMarker
	}

	err := ast.Walk(node, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		if entering {
			return ast.WalkContinue, nil
		}

		var (
			inMarker bool
			toRemove []ast.Node
		)

		for child := node.FirstChild(); child != nil; child = child.NextSibling() {
			if isMarker(child, source, startMarker) {
				inMarker = true
				toRemove = append(toRemove, child)
			}

			if isMarker(child, source, endMarker) {
				inMarker = false
				toRemove = append(toRemove, child)
			}

			codeBlock, ok := child.(*ast.CodeBlock)
			if !ok || codeBlock.CodeBlockKind != ast.CodeBlockKindFenced {
				continue
			}

			if inMarker {
				codeBlock.SetAttribute("data-killercoda-"+t.Kind, text.NewMultiLineValueFromString("true", text.IdentityDecoder))
				continue
			}

			if t.Kind == "exec" {
				continue
			}

			language, ok := codeBlock.Language(source)
			if ok && language == "bash" {
				codeBlock.SetAttribute("data-killercoda-exec", text.NewMultiLineValueFromString("true", text.IdentityDecoder))
			} else {
				codeBlock.SetAttribute("data-killercoda-copy", text.NewMultiLineValueFromString("true", text.IdentityDecoder))
			}
		}

		for _, child := range toRemove {
			node.RemoveChild(child)
		}

		return ast.WalkContinue, nil
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error transforming AST: %v\n", err)
	}
}

func parseShortcodeArgs(text []byte) map[string]string {
	args := make(map[string]string)
	for _, match := range kvPairRegexp.FindAllSubmatch(text, -1) {
		args[string(match[1])] = string(match[2])
	}

	return args
}

func isFigureShortcode(node ast.Node, source []byte) bool {
	if paragraph, ok := node.(*ast.Paragraph); ok {
		raw := bytes.TrimSpace(rawText(paragraph, source))

		return bytes.HasPrefix(raw, []byte("{{<")) && bytes.HasSuffix(raw, []byte(">}}")) && bytes.Contains(raw, []byte("figure"))
	}

	return false
}

func imageFromFigure(args map[string]string) *ast.Paragraph {
	var altText string

	if caption, ok := args["caption"]; ok {
		altText = caption
	}

	if alt, ok := args["alt"]; ok {
		altText = alt
	}

	paragraph := ast.NewParagraph()
	image := ast.NewImage(text.NewSingleLineValueFromString(args["src"], text.IdentityDecoder))
	image.AppendChild(ast.NewText(text.NewSingleLineValueFromString(altText, text.IdentityDecoder)))
	paragraph.AppendChild(image)

	return paragraph
}

type FigureTransformer struct{}

// Transform implements the parser.ASTTransformer interface and replaces all figure shortcodes with image elements.
func (t *FigureTransformer) Transform(node *ast.Document, reader text.Reader, _ parser.Context) {
	source := reader.Source()

	err := ast.Walk(node, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		replacements := map[ast.Node]ast.Node{}
		if !entering {
			return ast.WalkContinue, nil
		}

		for child := node.FirstChild(); child != nil; child = child.NextSibling() {
			if isFigureShortcode(child, source) {
				args := parseShortcodeArgs(bytes.TrimSpace(rawText(child, source)))
				replacement := imageFromFigure(args)

				replacements[child] = replacement
			}
		}

		for child, replacement := range replacements {
			node.ReplaceChild(child, replacement)
		}

		return ast.WalkContinue, nil
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error transforming AST: %v\n", err)
	}
}

type HeadingTransformer struct{}

// Transform implements the parser.ASTTransformer interface and ensures the heading hierarchy begins at H1.
func (t *HeadingTransformer) Transform(node *ast.Document, _ text.Reader, _ parser.Context) {
	var (
		headingDecrement  int
		foundFirstHeading bool
	)

	err := ast.Walk(node, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}

		if heading, ok := node.(*ast.Heading); ok {
			if !foundFirstHeading {
				foundFirstHeading = true

				headingDecrement = heading.Level - 1
			}

			heading.Level -= headingDecrement
		}

		return ast.WalkContinue, nil
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error transforming AST: %v\n", err)
	}
}

type IgnoreTransformer struct{}

// Transform implements the parser.ASTTransformer interface and removes all nodes between the ignore start and end markers.
func (t *IgnoreTransformer) Transform(node *ast.Document, reader text.Reader, _ parser.Context) {
	source := reader.Source()

	err := ast.Walk(node, func(node ast.Node, _ bool) (ast.WalkStatus, error) {
		var (
			inMarker bool
			toRemove []ast.Node
		)

		for child := node.FirstChild(); child != nil; child = child.NextSibling() {
			if isMarker(child, source, ignoreStartMarker) {
				inMarker = true
			}

			if inMarker {
				toRemove = append(toRemove, child)
			}

			if isMarker(child, source, ignoreEndMarker) {
				inMarker = false
			}
		}

		for _, child := range toRemove {
			node.RemoveChild(child)
		}

		return ast.WalkContinue, nil
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error transforming AST: %v\n", err)
	}
}

type InlineActionTransformer struct{}

// Transform implements the parser.ASTTransformer interface and adds inlineAction metadata to any fenced code blocks within between the start and end markers.
func (t *InlineActionTransformer) Transform(node *ast.Document, _ text.Reader, _ parser.Context) {
	err := ast.Walk(node, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		if entering {
			return ast.WalkContinue, nil
		}

		if node, ok := node.(*ast.CodeSpan); ok {
			node.SetAttribute("data-killercoda-copy", text.NewMultiLineValueFromString("true", text.IdentityDecoder))
		}

		return ast.WalkContinue, nil
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error transforming AST: %v\n", err)
	}
}

type LinkTransformer struct{}

// Transform implements the parser.ASTTransformer interface and replaces version substitution syntax (<SOMETHING_VERSION>) with 'latest' in links.
func (t *LinkTransformer) Transform(root *ast.Document, reader text.Reader, _ parser.Context) {
	source := reader.Source()

	err := ast.Walk(root, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}

		switch node := node.(type) {
		case *ast.Image:
			u, err := url.Parse(node.Destination.Str(source))
			if err != nil {
				return ast.WalkStop, fmt.Errorf("failed to parse URL: %w", err)
			}

			if u.Host == "" && u.Scheme == "" {
				u.Scheme = "https"
				u.Host = "grafana.com"
			}

			node.Destination = text.NewSingleLineValueFromString(u.String(), text.IdentityDecoder)
		case *ast.Link:
			replaced := versionSubstitutionRegexp.ReplaceAll(node.Destination.Bytes(source), []byte("latest"))
			u, err := url.Parse(string(replaced))
			if err != nil {
				return ast.WalkStop, fmt.Errorf("failed to parse URL: %w", err)
			}

			if u.Host == "" && u.Scheme == "" {
				u.Scheme = "https"
				u.Host = "grafana.com"
			}

			if u.Hostname() == "localhost" {
				destination := "{{TRAFFIC_HOST1_" + u.Port() + "}}"
				if u.Path != "" {
					destination += u.Path
				}
				node.Destination = text.NewSingleLineValueFromString(destination, text.IdentityDecoder)
			} else {
				node.Destination = text.NewSingleLineValueFromString(u.String(), text.IdentityDecoder)
			}
		}

		return ast.WalkContinue, nil
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error transforming AST: %v\n", err)
	}
}

type StepTransformer struct {
	StartMarker string
	EndMarker   string
}

// Transform implements the parser.ASTTransformer interface and keeps only the sibling nodes within the step start and end markers.
// It removes all other nodes resulting in a document that only contains the content between the markers.
// It removes the markers themselves.
func (t *StepTransformer) Transform(root *ast.Document, reader text.Reader, _ parser.Context) {
	source := reader.Source()

	err := ast.Walk(root, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}

		if isMarker(node, source, t.StartMarker) {
			var toKeep []ast.Node
			for sibling := node.NextSibling(); ; sibling = sibling.NextSibling() {
				if sibling == nil {
					return ast.WalkStop, fmt.Errorf("%w: %s", errNoEndMarker, t.StartMarker)
				}

				if isMarker(sibling, source, t.EndMarker) {
					break
				}

				toKeep = append(toKeep, sibling)
			}

			root.RemoveChildren()
			for _, keep := range toKeep {
				root.AppendChild(keep)
			}

			return ast.WalkStop, nil
		}

		return ast.WalkContinue, nil
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error transforming AST: %v\n", err)
	}
}
