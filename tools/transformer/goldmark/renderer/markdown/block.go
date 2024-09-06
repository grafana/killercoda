package markdown

import (
	"strings"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/util"
)

func (r *Renderer) renderBlockquote(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
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

func (r *Renderer) renderCodeBlock(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	const indent = "    "

	n := node.(*ast.CodeBlock)

	if entering {
		r.write(w, indent)
		r.pushPrefix(indent)
		r.writeLines(w, source, n)
	} else {
		r.popPrefix(indent)

		if node.NextSibling() != nil {
			r.write(w, '\n')
		}
	}

	return ast.WalkContinue, nil
}

func (r *Renderer) renderDocument(_ util.BufWriter, _ []byte, _ ast.Node, _ bool) (ast.WalkStatus, error) {
	return ast.WalkContinue, nil
}

func (r *Renderer) renderFencedCodeBlock(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.FencedCodeBlock)

	if entering {
		r.write(w, "```")

		language := n.Language(source)
		if language != nil {
			r.write(w, language)
		}

		r.write(w, '\n')
		r.writeLines(w, source, n)
	} else {
		r.write(w, "```")

		if r.Config.KillercodaActions {
			var action string

			if _, ok := n.AttributeString("data-killercoda-exec"); ok {
				action = "{{exec}}"
			} else if _, ok := n.AttributeString("data-killercoda-copy"); ok {
				action = "{{copy}}"
			}

			r.write(w, action)
		}

		r.write(w, '\n')

		if node.NextSibling() != nil {
			r.write(w, '\n')
		}
	}

	return ast.WalkContinue, nil
}

func (r *Renderer) renderHeading(w util.BufWriter, _ []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
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

func (r *Renderer) renderHTMLBlock(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.HTMLBlock)

	if entering {
		r.write(w, "<!-- raw HTML omitted -->\n")
	} else {
		if n.HasClosure() {
			r.write(w, "<!-- raw HTML omitted -->\n")
		}

		if n.NextSibling() != nil {
			r.write(w, '\n')
		}
	}

	return ast.WalkContinue, nil
}

func (r *Renderer) renderList(w util.BufWriter, _ []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		if node.NextSibling() != nil {
			r.write(w, '\n')
		}
	}

	return ast.WalkContinue, nil
}

func (r *Renderer) renderListItem(w util.BufWriter, _ []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	marker, indent := "- ", "  "
	if node.Parent().(*ast.List).IsOrdered() {
		marker, indent = "1. ", "   "
	}

	if entering {
		r.write(w, marker)
		r.pushPrefix(indent)
	} else {
		r.popPrefix(indent)

		if node.NextSibling() != nil {
			r.write(w, '\n')
		}
	}

	return ast.WalkContinue, nil
}

func (r *Renderer) renderParagraph(w util.BufWriter, _ []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		r.write(w, '\n')

		if node.NextSibling() != nil {
			r.write(w, '\n')
		}
	}

	return ast.WalkContinue, nil
}

func (r *Renderer) renderTextBlock(w util.BufWriter, _ []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		r.write(w, '\n')
	}

	return ast.WalkContinue, nil
}

func (r *Renderer) renderThematicBreak(w util.BufWriter, _ []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		return ast.WalkContinue, nil
	}

	r.write(w, "---\n")

	if node.NextSibling() != nil {
		r.write(w, '\n')
	}

	return ast.WalkContinue, nil
}
