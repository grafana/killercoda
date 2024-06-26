package main

import (
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

var (
	errNoEndMarker            = fmt.Errorf("no matching end marker found for start marker")
	versionSubstitutionRegexp = regexp.MustCompile(`<.+VERSION>`)
	kvPairRegexp              = regexp.MustCompile(`([a-zA-Z_]+)=(?:"(.*?[^\\])"|([^ ].+?))`)
)

const (
	copyStartMarker    = "<!-- INTERACTIVE copy START -->"
	copyEndMarker      = "<!-- INTERACTIVE copy END -->"
	execStartMarker    = "<!-- INTERACTIVE exec START -->"
	execEndMarker      = "<!-- INTERACTIVE exec END -->"
	finishStartMarker  = "<!-- INTERACTIVE finish.md START -->"
	finishEndMarker    = "<!-- INTERACTIVE finish.md END -->"
	ignoreStartMarker  = "<!-- INTERACTIVE ignore START -->"
	ignoreEndMarker    = "<!-- INTERACTIVE ignore END -->"
	includeStartMarker = "<!-- INTERACTIVE include START -->"
	includeEndMarker   = "<!-- INTERACTIVE include END -->"
	introStartMarker   = "<!-- INTERACTIVE intro.md START -->"
	introEndMarker     = "<!-- INTERACTIVE intro.md END -->"
)

func isMarker(node ast.Node, source []byte, marker string) bool {
	if node, ok := node.(*ast.HTMLBlock); ok {
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

			if inMarker {
				if fenced, ok := child.(*ast.FencedCodeBlock); ok {
					fenced.SetAttributeString("data-killercoda-"+t.Kind, "true")
				}
			}

			if isMarker(child, source, endMarker) {
				inMarker = false
				toRemove = append(toRemove, child)
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

type FinishTransformer struct{}

// Transform implements the parser.ASTTransformer interface and keeps only the sibling nodes within the finish start and end markers.
// It removes all other nodes resulting in a document that only contains the content between the markers.
// It removes the markers themselves.
func (t *FinishTransformer) Transform(root *ast.Document, reader text.Reader, _ parser.Context) {
	source := reader.Source()

	err := ast.Walk(root, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}

		if isMarker(node, source, finishStartMarker) {
			var toKeep []ast.Node
			for sibling := node.NextSibling(); ; sibling = sibling.NextSibling() {
				if sibling == nil {
					return ast.WalkStop, fmt.Errorf("%w: %s", errNoEndMarker, finishEndMarker)
				}

				if isMarker(sibling, source, finishEndMarker) {
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

type IncludeTransformer struct{}

// Transform implements the parser.ASTTransformer interface and keeps the Markdown content in HTML comments between the include start and end markers.
func (t *IncludeTransformer) Transform(node *ast.Document, reader text.Reader, _ parser.Context) {
	source := reader.Source()

	err := ast.Walk(node, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		var (
			inMarker     bool
			toRemove     []ast.Node
			replacements = make(map[ast.Node]ast.Node)
		)

		replacement := ast.NewParagraph()
		for child := node.FirstChild(); child != nil; child = child.NextSibling() {
			if isMarker(child, source, includeStartMarker) {
				inMarker = true
				toRemove = append(toRemove, child)

				continue
			}

			if isMarker(child, source, includeEndMarker) {
				inMarker = false

				replacements[child] = replacement
				replacement = ast.NewParagraph()

				continue
			}

			if inMarker {
				if comment, ok := child.(*ast.HTMLBlock); ok {
					text := rawText(comment, source)
					text = strings.TrimSpace(text)
					text = strings.TrimPrefix(text, "<!--")
					text = strings.TrimSuffix(text, "-->")
					text = strings.TrimSpace(text)

					str := ast.NewString([]byte(text + "\n"))
					str.SetRaw(true)

					replacement.AppendChild(replacement, str)
					toRemove = append(toRemove, child)
				}
			}
		}

		for _, child := range toRemove {
			node.RemoveChild(node, child)
		}

		for child, replacement := range replacements {
			node.ReplaceChild(node, child, replacement)
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

type IntroTransformer struct{}

// Transform implements the parser.ASTTransformer interface and keeps only the sibling nodes within the intro start and end markers.
// It removes all other nodes resulting in a document that only contains the content between the markers.
// It removes the markers themselves.
func (t *IntroTransformer) Transform(root *ast.Document, reader text.Reader, _ parser.Context) {
	source := reader.Source()

	err := ast.Walk(root, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}

		if isMarker(node, source, introStartMarker) {
			var toKeep []ast.Node
			for sibling := node.NextSibling(); ; sibling = sibling.NextSibling() {
				if sibling == nil {
					return ast.WalkStop, fmt.Errorf("%w: %s", errNoEndMarker, introStartMarker)
				}

				if isMarker(sibling, source, introEndMarker) {
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

type LinkTransformer struct{}

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
	Step int
}

// Transform implements the parser.ASTTransformer interface and keeps only the sibling nodes within the step start and end markers.
// It removes all other nodes resulting in a document that only contains the content between the markers.
// It removes the markers themselves.
func (t *StepTransformer) Transform(root *ast.Document, reader text.Reader, _ parser.Context) {
	stepStartMarker := fmt.Sprintf("<!-- INTERACTIVE step%d.md START -->", t.Step)
	stepEndMarker := fmt.Sprintf("<!-- INTERACTIVE step%d.md END -->", t.Step)

	source := reader.Source()

	err := ast.Walk(root, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}

		if isMarker(node, source, stepStartMarker) {
			var toKeep []ast.Node
			for sibling := node.NextSibling(); ; sibling = sibling.NextSibling() {
				if sibling == nil {
					return ast.WalkStop, fmt.Errorf("%w: %s", errNoEndMarker, stepStartMarker)
				}

				if isMarker(sibling, source, stepEndMarker) {
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
