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
)

const (
	copyStartMarker   = "<!-- Killercoda copy START -->"
	copyEndMarker     = "<!-- Killercoda copy END -->"
	execStartMarker   = "<!-- Killercoda exec START -->"
	execEndMarker     = "<!-- Killercoda exec END -->"
	ignoreStartMarker = "<!-- Killercoda ignore START -->"
	ignoreEndMarker   = "<!-- Killercoda ignore END -->"
	introStartMarker  = "<!-- Killercoda intro.md START -->"
	introEndMarker    = "<!-- Killercoda intro.md END -->"
	stepStartMarker   = "<!-- Killercoda step1.md START -->"
	stepEndMarker     = "<!-- Killercoda step1.md END -->"
)

func isMarker(node ast.Node, source []byte, marker string) bool {
	if node, ok := node.(*ast.HTMLBlock); ok {
		l := node.Lines().Len()
		for i := 0; i < l; i++ {
			line := node.Lines().At(i)
			if strings.TrimSpace(string(line.Value(source))) == marker {
				return true
			}
		}
	}

	return false
}

type CopyTransformer struct{}

// Transform implements the parser.ASTTransformer interface and adds copy metadata to any fenced code blocks within between the copy start and end markers.
func (t *CopyTransformer) Transform(node *ast.Document, reader text.Reader, _ parser.Context) {
	source := reader.Source()

	err := ast.Walk(node, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}

		if isMarker(node, source, copyStartMarker) {
			toRemove := []ast.Node{
				node,
			}

			for sibling := node.NextSibling(); ; sibling = sibling.NextSibling() {
				if sibling == nil {
					return ast.WalkStop, fmt.Errorf("%w: %s", errNoEndMarker, ignoreStartMarker)
				}

				if fenced, ok := sibling.(*ast.FencedCodeBlock); ok {
					fenced.SetAttributeString("data-killercoda-copy", "true")
				}

				if isMarker(sibling, source, copyEndMarker) {
					toRemove = append(toRemove, sibling)

					break
				}
			}

			for _, node := range toRemove {
				node.Parent().RemoveChild(node.Parent(), node)
			}
		}

		return ast.WalkContinue, nil
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error transforming AST: %v\n", err)
	}
}

type ExecTransformer struct{}

// Transform implements the parser.ASTTransformer interface and adds exec metadata to any fenced code blocks within between the exec start and end markers.
func (t *ExecTransformer) Transform(node *ast.Document, reader text.Reader, _ parser.Context) {
	source := reader.Source()

	err := ast.Walk(node, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}

		if isMarker(node, source, execStartMarker) {
			toRemove := []ast.Node{
				node,
			}

			for sibling := node.NextSibling(); ; sibling = sibling.NextSibling() {
				if sibling == nil {
					return ast.WalkStop, fmt.Errorf("%w: %s", errNoEndMarker, ignoreStartMarker)
				}

				if fenced, ok := sibling.(*ast.FencedCodeBlock); ok {
					fenced.SetAttributeString("data-killercoda-exec", "true")
				}

				if isMarker(sibling, source, execEndMarker) {
					toRemove = append(toRemove, sibling)

					break
				}
			}

			for _, node := range toRemove {
				node.Parent().RemoveChild(node.Parent(), node)
			}
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
		if !entering {
			return ast.WalkContinue, nil
		}

		if isMarker(node, source, ignoreStartMarker) {
			toRemove := []ast.Node{
				node,
			}

			for sibling := node.NextSibling(); ; sibling = sibling.NextSibling() {
				if sibling == nil {
					return ast.WalkStop, fmt.Errorf("%w: %s", errNoEndMarker, ignoreStartMarker)
				}

				toRemove = append(toRemove, sibling)

				if isMarker(sibling, source, ignoreEndMarker) {
					break
				}
			}

			for _, node := range toRemove {
				node.Parent().RemoveChild(node.Parent(), node)
			}
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
func (t *IntroTransformer) Transform(root *ast.Document, reader text.Reader, pc parser.Context) {
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
func (t *LinkTransformer) Transform(root *ast.Document, reader text.Reader, pc parser.Context) {
	err := ast.Walk(root, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}

		if link, ok := node.(*ast.Link); ok {
			link.Destination = versionSubstitutionRegexp.ReplaceAll(link.Destination, []byte("latest"))

			u, err := url.Parse(string(link.Destination))
			if err != nil {
				return ast.WalkStop, fmt.Errorf("failed to parse URL: %w", err)
			}

			if u.Host == "" && u.Scheme == "" {
				u.Scheme = "https"
				u.Host = "grafana.com"
			}

			if u.Hostname() == "localhost" {
				link.Destination = []byte("{{TRAFFIC_HOST1_" + u.Port() + "}}")
			} else {
				link.Destination = []byte(u.String())
			}
		}

		return ast.WalkContinue, nil
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error transforming AST: %v\n", err)
	}
}

type StepTransformer struct{}

// Transform implements the parser.ASTTransformer interface and keeps only the sibling nodes within the step start and end markers.
// It removes all other nodes resulting in a document that only contains the content between the markers.
// It removes the markers themselves.
func (t *StepTransformer) Transform(root *ast.Document, reader text.Reader, _ parser.Context) {
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
