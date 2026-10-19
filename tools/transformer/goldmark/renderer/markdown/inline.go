package markdown

import (
	"bytes"

	"github.com/yuin/goldmark/v2/ast"
	"github.com/yuin/goldmark/v2/renderer"
	"github.com/yuin/goldmark/v2/util"
)

func (r *Renderer) renderAutoLink(w util.BufWriter, source []byte, node ast.Node, entering bool, _ renderer.Context) (ast.WalkStatus, error) {
	if !entering {
		return ast.WalkContinue, nil
	}

	n := node.(*ast.AutoLink)
	if !n.Text.IsEmpty() {
		r.write(w, n.Text.Bytes(source))
		return ast.WalkContinue, nil
	}

	r.write(w, '<')
	r.write(w, util.EscapeHTML(util.URLEscape(n.Destination.Bytes(source))))
	r.write(w, '>')

	return ast.WalkContinue, nil
}

func (r *Renderer) renderCodeSpan(w util.BufWriter, source []byte, node ast.Node, entering bool, _ renderer.Context) (ast.WalkStatus, error) {
	n := node.(*ast.CodeSpan)

	if entering {
		r.write(w, '`')
		value := n.Value.Bytes(source)
		if bytes.HasSuffix(value, []byte("\n")) {
			r.write(w, value[:len(value)-1])
			r.write(w, []byte(" "))
		} else {
			r.write(w, value)
		}
		return ast.WalkSkipChildren, nil
	}

	r.write(w, '`')

	if r.Config().KillercodaActions {
		if _, ok := node.Attribute("data-killercoda-copy"); ok {
			r.write(w, "{{copy}}")
		}
	}

	return ast.WalkContinue, nil
}

// renderEmphasis renders emphasis.
// Correctly rendering emphasis is a lot more complicated that this function.
// https://spec.commonmark.org/0.31.2/#emphasis-and-strong-emphasis
func (r *Renderer) renderEmphasis(w util.BufWriter, _ []byte, _ ast.Node, _ bool, _ renderer.Context) (ast.WalkStatus, error) {
	r.write(w, "_")

	return ast.WalkContinue, nil
}

func (r *Renderer) renderStrong(w util.BufWriter, _ []byte, _ ast.Node, _ bool, _ renderer.Context) (ast.WalkStatus, error) {
	r.write(w, "**")

	return ast.WalkContinue, nil
}

func (r *Renderer) renderImage(w util.BufWriter, source []byte, node ast.Node, entering bool, _ renderer.Context) (ast.WalkStatus, error) {
	if entering {
		r.write(w, "![")
	} else {
		n := node.(*ast.Image)

		r.write(w, "](")
		r.write(w, n.Destination.Bytes(source))

		if !n.Title.IsEmpty() {
			r.write(w, " \"")
			r.write(w, n.Title.Bytes(source))
			r.write(w, "\"")
		}

		r.write(w, ')')
	}

	return ast.WalkContinue, nil
}

func (r *Renderer) renderLink(w util.BufWriter, source []byte, node ast.Node, entering bool, _ renderer.Context) (ast.WalkStatus, error) {
	n := node.(*ast.Link)

	if entering {
		r.write(w, '[')
	} else {
		r.write(w, "](")
		r.write(w, n.Destination.Bytes(source))

		if !n.Title.IsEmpty() {
			r.write(w, " \"")
			r.write(w, n.Title.Bytes(source))
			r.write(w, "\"")
		}

		r.write(w, ')')
	}

	return ast.WalkContinue, nil
}

func (r *Renderer) renderRawHTML(w util.BufWriter, source []byte, node ast.Node, entering bool, _ renderer.Context) (ast.WalkStatus, error) {
	if !entering {
		return ast.WalkSkipChildren, nil
	}

	if r.Config().Unsafe {
		n := node.(*ast.RawHTML)
		r.secureWrite(w, n.Value.Bytes(source))
	} else {
		r.write(w, "<!-- raw HTML omitted -->")
	}

	return ast.WalkSkipChildren, nil
}

func (r *Renderer) renderText(w util.BufWriter, source []byte, node ast.Node, entering bool, _ renderer.Context) (ast.WalkStatus, error) {
	if !entering {
		return ast.WalkContinue, nil
	}

	n := node.(*ast.Text)
	r.write(w, n.Value.Bytes(source))

	if n.HardLineBreak() {
		r.write(w, "\n\n")
	} else if n.SoftLineBreak() {
		r.write(w, '\n')
	}

	return ast.WalkContinue, nil
}
