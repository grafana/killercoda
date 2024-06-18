package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

const (
	ignoreStartMarker = "<!-- Killercoda ignore START -->"
	ignoreEndMarker   = "<!-- Killercoda ignore END -->"
)

type IgnoreTransformer struct{}

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
					return ast.WalkStop, fmt.Errorf("no matching end marker found for ignore start marker")
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
