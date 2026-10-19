// Package extension implements Goldmark extensions.
package extension

import (
	gast "github.com/yuin/goldmark/v2/ast"
	"github.com/yuin/goldmark/v2/extension/ast"
	"github.com/yuin/goldmark/v2/renderer"
	"github.com/yuin/goldmark/v2/util"
)

// TableNodeRenderers returns node renderers for rendering Goldmark table nodes as Markdown.
func TableNodeRenderers() map[gast.NodeKind]renderer.NodeRenderer[util.BufWriter] {
	r := &tableMarkdownRenderer{}

	return map[gast.NodeKind]renderer.NodeRenderer[util.BufWriter]{
		ast.KindTable:       renderer.NodeRendererFunc(r.renderTable),
		ast.KindTableHeader: renderer.NodeRendererFunc(r.renderTableHeader),
		ast.KindTableRow:    renderer.NodeRendererFunc(r.renderTableRow),
		ast.KindTableCell:   renderer.NodeRendererFunc(r.renderTableCell),
	}
}

type tableMarkdownRenderer struct{}

func (r *tableMarkdownRenderer) renderTable(w util.BufWriter, _ []byte, _ gast.Node, entering bool, _ renderer.Context) (gast.WalkStatus, error) {
	if !entering {
		_, _ = w.WriteString("\n")
	}

	return gast.WalkContinue, nil
}

func (r *tableMarkdownRenderer) renderTableHeader(w util.BufWriter, _ []byte, n gast.Node, entering bool, _ renderer.Context) (gast.WalkStatus, error) {
	if !entering {
		_, _ = w.WriteString("\n")

		_, _ = w.WriteString("|")
		for i := 0; i < n.ChildCount(); i++ {
			_, _ = w.WriteString(" --- |")
		}
		_, _ = w.WriteString("\n")
	}

	return gast.WalkContinue, nil
}

func (r *tableMarkdownRenderer) renderTableRow(w util.BufWriter, _ []byte, _ gast.Node, entering bool, _ renderer.Context) (gast.WalkStatus, error) {
	if !entering {
		_, _ = w.WriteString("\n")
	}

	return gast.WalkContinue, nil
}

func (r *tableMarkdownRenderer) renderTableCell(w util.BufWriter, _ []byte, n gast.Node, entering bool, _ renderer.Context) (gast.WalkStatus, error) {
	if entering {
		_, _ = w.WriteString("| ")
	} else if n.NextSibling() != nil {
		_, _ = w.WriteString(" ")
	} else {
		_, _ = w.WriteString(" |")
	}

	return gast.WalkContinue, nil
}
