package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

var errNoEndMarker = fmt.Errorf("no matching end marker found for start marker")

const (
	ignoreStartMarker = "<!-- Killercoda ignore START -->"
	ignoreEndMarker   = "<!-- Killercoda ignore END -->"
	introStartMarker  = "<!-- Killercoda intro.md START -->"
	introEndMarker    = "<!-- Killercoda intro.md END -->"
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

type IgnoreTransformer struct{}

// Transform implements the parser.ASTTransformer interface and removes all nodes between the ignore start and end markers.
func (t *IgnoreTransformer) Transform(node *ast.Document, reader text.Reader, pc parser.Context) {
	source := reader.Source()

	err := ast.Walk(node, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}

		if isMarker(node, source, ignoreStartMarker) {
			var end bool
			toRemove := []ast.Node{
				node,
			}

			for sibling := node.NextSibling(); !end; sibling = sibling.NextSibling() {
				if sibling == nil {
					return ast.WalkStop, fmt.Errorf("%w: %s", errNoEndMarker, ignoreStartMarker)
				}

				if isMarker(sibling, source, ignoreEndMarker) {
					end = true
				}

				toRemove = append(toRemove, sibling)
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
