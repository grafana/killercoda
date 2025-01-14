// Package extension implements Goldmark extensions.
package extension

import (
	"github.com/yuin/goldmark"
	gast "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/extension/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

// TableMarkdownRenderer is a renderer.NodeRenderer implementation that
// renders Table nodes.
type TableMarkdownRenderer struct {
	extension.TableConfig
}

// NewTableMarkdownRenderer returns a new TableMarkdownRenderer.
func NewTableMarkdownRenderer(opts ...extension.TableOption) renderer.NodeRenderer {
	r := &TableMarkdownRenderer{
		TableConfig: extension.NewTableConfig(),
	}
	for _, opt := range opts {
		opt.SetTableOption(&r.TableConfig)
	}
	return r
}

// RegisterFuncs implements renderer.NodeRenderer.RegisterFuncs.
func (r *TableMarkdownRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(ast.KindTable, r.renderTable)
	reg.Register(ast.KindTableHeader, r.renderTableHeader)
	reg.Register(ast.KindTableRow, r.renderTableRow)
	reg.Register(ast.KindTableCell, r.renderTableCell)
}

func (r *TableMarkdownRenderer) renderTable(w util.BufWriter, source []byte, n gast.Node, entering bool) (gast.WalkStatus, error) {
	if !entering {
		_, _ = w.WriteString("\n")
	}

	return gast.WalkContinue, nil
}

func (r *TableMarkdownRenderer) renderTableHeader(w util.BufWriter, source []byte, n gast.Node, entering bool) (gast.WalkStatus, error) {
	if !entering {
		_, _ = w.WriteString("\n")

		// Add a separator line after the header row.
		_, _ = w.WriteString("|")
		for i := 0; i < n.ChildCount(); i++ {
			_, _ = w.WriteString(" --- |")
		}
		_, _ = w.WriteString("\n")
	}

	return gast.WalkContinue, nil
}

func (r *TableMarkdownRenderer) renderTableRow(w util.BufWriter, _ []byte, n gast.Node, entering bool) (gast.WalkStatus, error) {
	if !entering {
		_, _ = w.WriteString("\n")
	}

	return gast.WalkContinue, nil
}

func (r *TableMarkdownRenderer) renderTableCell(w util.BufWriter, _ []byte, n gast.Node, entering bool) (gast.WalkStatus, error) {
	if entering {
		_, _ = w.WriteString("| ")
	} else if n.NextSibling() != nil {
		_, _ = w.WriteString(" ")
	} else {
		_, _ = w.WriteString(" |")
	}

	return gast.WalkContinue, nil
}

type table struct {
	options []extension.TableOption
}

// Table is an extension that allow you to use GFM tables .
var Table = &table{
	options: []extension.TableOption{},
}

// NewTable returns a new extension with given options.
func NewTable(opts ...extension.TableOption) goldmark.Extender {
	return &table{
		options: opts,
	}
}

func (e *table) Extend(m goldmark.Markdown) {
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(NewTableMarkdownRenderer(e.options...), 500),
	))
}
