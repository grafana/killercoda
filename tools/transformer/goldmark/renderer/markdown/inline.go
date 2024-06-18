package markdown

import (
	"bytes"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/util"
)

func (r *Renderer) renderCodeSpan(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		r.Write(w, '`')

		for c := node.FirstChild(); c != nil; c = c.NextSibling() {
			segment := c.(*ast.Text).Segment
			value := segment.Value(source)

			if bytes.HasSuffix(value, []byte("\n")) {
				r.Writer.RawWrite(w, value[:len(value)-1])
				r.Writer.RawWrite(w, []byte(" "))
			} else {
				r.Writer.RawWrite(w, value)
			}
		}

		return ast.WalkSkipChildren, nil
	}

	r.Write(w, '`')

	return ast.WalkContinue, nil
}

func (r *Renderer) renderEmphasis(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.Emphasis)

	delim := "_"
	if n.Level == 2 {
		delim = "**"
	}

	if entering {
		r.Write(w, delim)
	} else {
		r.Write(w, delim)
	}

	return ast.WalkContinue, nil
}

func (r *Renderer) renderLink(w util.BufWriter, _ []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.Link)

	if entering {
		r.Write(w, '[')
	} else {
		r.Write(w, "](")
		r.Write(w, n.Destination)

		if n.Title != nil {
			r.Write(w, " \"")
			r.Write(w, n.Title)
			r.Write(w, "\"")
		}
		r.Write(w, ')')
	}

	return ast.WalkContinue, nil
}

func (r *Renderer) renderText(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		return ast.WalkContinue, nil
	}

	n := node.(*ast.Text)
	segment := n.Segment
	if n.IsRaw() {
		panic("TODO: implement")
	} else {
		value := segment.Value(source)
		r.Write(w, value)
		if n.HardLineBreak() {
			r.Write(w, "\n\n")
		} else if n.SoftLineBreak() {
			r.Write(w, '\n')
		}
	}

	return ast.WalkContinue, nil
}
