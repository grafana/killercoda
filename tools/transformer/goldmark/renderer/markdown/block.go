package markdown

import (
	"strings"

	"github.com/yuin/goldmark/v2/ast"
	"github.com/yuin/goldmark/v2/renderer"
	"github.com/yuin/goldmark/v2/util"
)

func (r *Renderer) renderBlockquote(w util.BufWriter, _ []byte, node ast.Node, entering bool, _ renderer.Context) (ast.WalkStatus, error) {
	const prefix = "> "
	if entering {
		r.write(w, prefix)
		r.pushPrefix(prefix)
	} else {
		r.popPrefix(prefix)
		if node.NextSibling() != nil {
			r.write(w, '\n')
		}
	}

	return ast.WalkContinue, nil
}

func (r *Renderer) renderCodeBlock(w util.BufWriter, source []byte, node ast.Node, entering bool, _ renderer.Context) (ast.WalkStatus, error) {
	n := node.(*ast.CodeBlock)

	if entering {
		if n.CodeBlockKind == ast.CodeBlockKindFenced {
			r.write(w, "```")
			if language, ok := n.Language(source); ok {
				r.write(w, language)
			}
			r.write(w, '\n')
		} else {
			const indent = "    "
			r.write(w, indent)
			r.pushPrefix(indent)
		}

		for _, line := range n.Value.Segments() {
			r.write(w, line.Bytes(source))
		}
	} else {
		if n.CodeBlockKind == ast.CodeBlockKindFenced {
			r.write(w, "```")

			if r.Config().KillercodaActions {
				var action string
				if _, ok := n.Attribute("data-killercoda-exec"); ok {
					action = "{{exec}}"
				} else if _, ok := n.Attribute("data-killercoda-copy"); ok {
					action = "{{copy}}"
				}
				r.write(w, action)
			}
			r.write(w, '\n')
		} else {
			r.popPrefix("    ")
		}

		if node.NextSibling() != nil {
			r.write(w, '\n')
		}
	}

	return ast.WalkContinue, nil
}

func (r *Renderer) renderDocument(_ util.BufWriter, _ []byte, _ ast.Node, _ bool, _ renderer.Context) (ast.WalkStatus, error) {
	return ast.WalkContinue, nil
}

func (r *Renderer) renderHeading(w util.BufWriter, _ []byte, node ast.Node, entering bool, _ renderer.Context) (ast.WalkStatus, error) {
	n := node.(*ast.Heading)

	if entering {
		r.write(w, strings.Repeat("#", n.Level))
		r.write(w, ' ')
	} else {
		r.write(w, '\n')

		if node.NextSibling() != nil {
			r.write(w, '\n')
		}
	}

	return ast.WalkContinue, nil
}

func (r *Renderer) renderHTMLBlock(w util.BufWriter, source []byte, node ast.Node, entering bool, _ renderer.Context) (ast.WalkStatus, error) {
	if !entering {
		if node.NextSibling() != nil {
			r.write(w, '\n')
		}
		return ast.WalkContinue, nil
	}

	n := node.(*ast.HTMLBlock)
	if r.Config().Unsafe {
		for _, line := range n.Value.Segments() {
			r.secureWrite(w, line.Bytes(source))
		}
	} else {
		r.write(w, "<!-- raw HTML omitted -->\n")
	}

	return ast.WalkContinue, nil
}

func (r *Renderer) renderList(w util.BufWriter, _ []byte, node ast.Node, entering bool, _ renderer.Context) (ast.WalkStatus, error) {
	if !entering && node.NextSibling() != nil {
		r.write(w, '\n')
	}

	return ast.WalkContinue, nil
}

func (r *Renderer) renderListItem(w util.BufWriter, _ []byte, node ast.Node, entering bool, _ renderer.Context) (ast.WalkStatus, error) {
	marker, indent := "- ", "  "
	if node.Parent().(*ast.List).IsOrdered() {
		marker, indent = "1. ", "   "
	}

	if entering {
		r.write(w, marker)
		r.pushPrefix(indent)
	} else {
		r.popPrefix(indent)
	}

	return ast.WalkContinue, nil
}

func (r *Renderer) renderParagraph(w util.BufWriter, _ []byte, node ast.Node, entering bool, _ renderer.Context) (ast.WalkStatus, error) {
	if !entering {
		r.write(w, '\n')

		tightListParagraph := false
		if listItem, ok := node.Parent().(*ast.ListItem); ok {
			if list, ok := listItem.Parent().(*ast.List); ok {
				tightListParagraph = list.IsTight
			}
		}

		if node.NextSibling() != nil && !tightListParagraph {
			r.write(w, '\n')
		}
	}

	return ast.WalkContinue, nil
}

func (r *Renderer) renderTextBlock(w util.BufWriter, _ []byte, _ ast.Node, entering bool, _ renderer.Context) (ast.WalkStatus, error) {
	if !entering {
		r.write(w, '\n')
	}

	return ast.WalkContinue, nil
}

func (r *Renderer) renderThematicBreak(w util.BufWriter, _ []byte, node ast.Node, entering bool, _ renderer.Context) (ast.WalkStatus, error) {
	if !entering {
		return ast.WalkContinue, nil
	}

	r.write(w, "---\n")

	if node.NextSibling() != nil {
		r.write(w, '\n')
	}

	return ast.WalkContinue, nil
}
