package markdown

import (
	"strings"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/util"
)

func (r *Renderer) renderFencedCodeBlock(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.FencedCodeBlock)

	if entering {
		r.Write(w, "```")

		language := n.Language(source)
		if language != nil {
			r.Write(w, language)
		}

		r.Write(w, '\n')
		r.writeLines(w, source, n)
	} else {
		r.Write(w, "```\n")
	}

	return ast.WalkContinue, nil
}

func (r *Renderer) renderHeading(w util.BufWriter, _ []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.Heading)

	if entering {
		r.Write(w, strings.Repeat("#", n.Level))
		r.Write(w, ' ')
	} else {
		if r.Attributes {
			r.Write(w, " {")
			if n.Attributes() != nil {
				RenderAttributes(w, node, HeadingAttributeFilter)
			}
			r.Write(w, '}')
		}

		r.Write(w, '\n')

		if node.NextSibling() != nil {
			r.Write(w, '\n')
		}
	}

	return ast.WalkContinue, nil
}

func (r *Renderer) renderHTMLBlock(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.HTMLBlock)

	if entering {
		if r.Unsafe {
			l := n.Lines().Len()
			for i := 0; i < l; i++ {
				line := n.Lines().At(i)
				r.Writer.SecureWrite(w, line.Value(source))
			}
		} else {
			r.Write(w, "<!-- raw HTML omitted -->\n")
		}
	} else {
		if n.HasClosure() {
			if r.Unsafe {
				closure := n.ClosureLine
				r.Writer.SecureWrite(w, closure.Value(source))
			} else {
				r.Write(w, "<!-- raw HTML omitted -->\n")
			}
		}

		if n.NextSibling() != nil {
			r.Write(w, '\n')
		}
	}

	return ast.WalkContinue, nil
}

func (r *Renderer) renderList(w util.BufWriter, _ []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		if r.Attributes {
			if node.Attributes() != nil {
				r.Write(w, "\n{")
				RenderAttributes(w, node, ListAttributeFilter)
				r.Write(w, "}")
			}
		}

		if node.NextSibling() != nil {
			r.Write(w, '\n')
		}
	}

	return ast.WalkContinue, nil
}

func (r *Renderer) renderListItem(w util.BufWriter, _ []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	marker := "- "
	indent := 2
	if node.Parent().(*ast.List).IsOrdered() {
		marker = "1. "
		indent = 3
	}

	if entering {
		r.Write(w, marker)
		r.indent += indent
	} else {
		r.indent -= indent
	}

	return ast.WalkContinue, nil
}

func (r *Renderer) renderParagraph(w util.BufWriter, _ []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		if r.Attributes {
			r.Write(w, "\n{")
			RenderAttributes(w, node, ParagraphAttributeFilter)
			r.Write(w, "}")
		}

		r.Write(w, '\n')
		if node.NextSibling() != nil {
			r.Write(w, '\n')
		}
	}

	return ast.WalkContinue, nil
}

func (r *Renderer) renderTextBlock(w util.BufWriter, _ []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		r.Write(w, '\n')
	}

	return ast.WalkContinue, nil
}
