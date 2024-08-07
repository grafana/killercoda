package main

import (
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

var (
	errNoEndMarker            = fmt.Errorf("no matching end marker found for start marker")
	versionSubstitutionRegexp = regexp.MustCompile(`<.+VERSION>`)
	kvPairRegexp              = regexp.MustCompile(`([a-zA-Z_]+)=(?:"(.*?[^\\])"|([^ ].+?))`)
)

func isMarker(node ast.Node, source []byte, marker string) bool {
	switch node := node.(type) {
	case *ast.Text, *ast.String:
		if strings.TrimSpace(string(node.Text(source))) == marker {
			return true
		}

	case *ast.HTMLBlock, *ast.Paragraph:
		raw := rawText(node, source)
		if strings.TrimSpace(raw) == marker {
			return true
		}
	}

	return false
}

func rawText(node ast.Node, source []byte) string {
	builder := &strings.Builder{}

	if node.Type() == ast.TypeBlock {
		for i := 0; i < node.Lines().Len(); i++ {
			line := node.Lines().At(i)
			builder.Write(line.Value(source))
		}
	}

	return builder.String()
}

type ActionTransformer struct {
	Kind string
}

func (t *ActionTransformer) Extend(md goldmark.Markdown) {
	md.Parser().AddOptions(
		parser.WithASTTransformers(
			util.Prioritized(t, 0)))
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

			if fenced, ok := child.(*ast.FencedCodeBlock); ok {
				if inMarker {
					fenced.SetAttributeString("data-killercoda-"+t.Kind, "true")
				} else {
					// Only set the language attribute if not within a marker
					if t.Kind != "exec" {
						language := string(fenced.Language(source))
						if language == "bash" {
							fenced.SetAttributeString("data-killercoda-exec", "true")
						} else {
							fenced.SetAttributeString("data-killercoda-copy", "true")
						}
					}
				}
			}
		}

		for _, child := range toRemove {
			node.RemoveChild(node, child)
		}

		return ast.WalkContinue, nil
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error transforming AST: %v\n", err)
	}
}

type AdmonitionTransformer struct{}

func (t *AdmonitionTransformer) Extend(md goldmark.Markdown) {
	md.Parser().AddOptions(
		parser.WithASTTransformers(
			util.Prioritized(t, 0)))
}

// Transform implements the parser.ASTTransformer interface and replaces all admonition shortcodes with blockquotes.
func (t *AdmonitionTransformer) Transform(node *ast.Document, reader text.Reader, _ parser.Context) {
	source := reader.Source()

	err := ast.Walk(node, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}

		if paragraph, ok := node.(*ast.Paragraph); ok {
			raw := strings.TrimSpace(rawText(paragraph, source))

			if strings.HasPrefix(raw, "{{<") && strings.HasSuffix(raw, ">}}") && strings.Contains(raw, "admonition") {
				panic("TODO: implement")
			}
		}

		return ast.WalkContinue, nil
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error transforming AST: %v\n", err)
	}
}

func isFigureShortcode(node ast.Node, source []byte) bool {
	if paragraph, ok := node.(*ast.Paragraph); ok {
		raw := strings.TrimSpace(rawText(paragraph, source))

		return strings.HasPrefix(raw, "{{<") && strings.HasSuffix(raw, ">}}") && strings.Contains(raw, "figure")
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

	text := ast.NewString([]byte(altText))
	text.SetRaw(true)

	link := ast.NewLink()
	link.Destination = []byte(args["src"])
	link.AppendChild(link, text)

	paragraph := ast.NewParagraph()
	paragraph.AppendChild(paragraph, ast.NewImage(link))

	return paragraph
}

type FigureTransformer struct{}

func (t *FigureTransformer) Extend(md goldmark.Markdown) {
	md.Parser().AddOptions(
		parser.WithASTTransformers(
			util.Prioritized(t, 0)))
}

// Transform implements the parser.ASTTransformer interface and replaces all figure shortcodes with image elements.
func (t *FigureTransformer) Transform(node *ast.Document, reader text.Reader, _ parser.Context) {
	source := reader.Source()

	err := ast.Walk(node, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}

		if isFigureShortcode(node, source) {
			raw := strings.TrimSpace(rawText(node, source))

			args := make(map[string]string)
			for _, match := range kvPairRegexp.FindAllStringSubmatch(raw, -1) {
				args[match[1]] = match[2]
			}

			replacement := imageFromFigure(args)

			node.Parent().ReplaceChild(node.Parent(), node, replacement)
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

func (t *IgnoreTransformer) Extend(md goldmark.Markdown) {
	md.Parser().AddOptions(
		parser.WithASTTransformers(
			util.Prioritized(t, 0)))
}

// Transform implements the parser.ASTTransformer interface and removes all nodes between the ignore start and end markers.
func (t *IgnoreTransformer) Transform(node *ast.Document, reader text.Reader, _ parser.Context) {
	source := reader.Source()

	err := ast.Walk(node, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
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
			node.RemoveChild(node, child)
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
			node.SetAttributeString("data-killercoda-copy", "true")
		}

		return ast.WalkContinue, nil
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error transforming AST: %v\n", err)
	}
}

type LinkTransformer struct{}

func (t *LinkTransformer) Extend(md goldmark.Markdown) {
	md.Parser().AddOptions(
		parser.WithASTTransformers(
			util.Prioritized(t, 0)))
}

// Transform implements the parser.ASTTransformer interface and replaces version substitution syntax (<SOMETHING_VERSION>) with 'latest' in links.
func (t *LinkTransformer) Transform(root *ast.Document, _ text.Reader, _ parser.Context) {
	err := ast.Walk(root, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}

		switch node := node.(type) {
		case *ast.Image:
			u, err := url.Parse(string(node.Destination))
			if err != nil {
				return ast.WalkStop, fmt.Errorf("failed to parse URL: %w", err)
			}

			if u.Host == "" && u.Scheme == "" {
				u.Scheme = "https"
				u.Host = "grafana.com"
			}

			node.Destination = []byte(u.String())
		case *ast.Link:
			node.Destination = versionSubstitutionRegexp.ReplaceAll(node.Destination, []byte("latest"))

			u, err := url.Parse(string(node.Destination))
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
				node.Destination = []byte(destination)
			} else {
				node.Destination = []byte(u.String())
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

func (t *StepTransformer) Extend(md goldmark.Markdown) {
	md.Parser().AddOptions(
		parser.WithASTTransformers(
			util.Prioritized(t, 0)))
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

			root.RemoveChildren(root)

			for _, node := range toKeep {
				root.AppendChild(root, node)
				node.SetParent(root)
			}

			return ast.WalkStop, nil
		}

		return ast.WalkContinue, nil
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error transforming AST: %v\n", err)
	}
}
