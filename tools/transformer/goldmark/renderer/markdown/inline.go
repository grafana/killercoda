package markdown

import (
	"bytes"
	"html"

	"github.com/yuin/goldmark/ast"
	rendererHTML "github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"
)

func (r *Renderer) renderAutoLink(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	panic("TODO: implement")
	n := node.(*ast.AutoLink)
	if !entering {
		return ast.WalkContinue, nil
	}
	_, _ = w.WriteString(`<a href="`)
	url := n.URL(source)
	label := n.Label(source)
	if n.AutoLinkType == ast.AutoLinkEmail && !bytes.HasPrefix(bytes.ToLower(url), []byte("mailto:")) {
		_, _ = w.WriteString("mailto:")
	}
	_, _ = w.Write(util.EscapeHTML(util.URLEscape(url, false)))
	if n.Attributes() != nil {
		_ = w.WriteByte('"')
		rendererHTML.RenderAttributes(w, n, rendererHTML.LinkAttributeFilter)
		_ = w.WriteByte('>')
	} else {
		_, _ = w.WriteString(`">`)
	}
	_, _ = w.Write(util.EscapeHTML(label))
	_, _ = w.WriteString(`</a>`)
	return ast.WalkContinue, nil
}

func (r *Renderer) renderCodeSpan(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		r.write(w, '`')

		for c := node.FirstChild(); c != nil; c = c.NextSibling() {
			segment := c.(*ast.Text).Segment
			value := segment.Value(source)

			if bytes.HasSuffix(value, []byte("\n")) {
				r.write(w, value[:len(value)-1])
				r.write(w, []byte(" "))
			} else {
				r.write(w, value)
			}
		}

		return ast.WalkSkipChildren, nil
	}

	r.write(w, '`')

	return ast.WalkContinue, nil
}

func (r *Renderer) renderEmphasis(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.Emphasis)

	delim := "_"
	if n.Level == 2 {
		delim = "**"
	}

	if entering {
		r.write(w, delim)
	} else {
		r.write(w, delim)
	}

	return ast.WalkContinue, nil
}

func (r *Renderer) renderImage(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		r.write(w, "![")
	} else {
		n := node.(*ast.Image)

		r.write(w, "](")
		r.write(w, n.Destination)

		if n.Title != nil {
			r.write(w, " \"")
			r.write(w, n.Title)
			r.write(w, "\"")
		}
		r.write(w, ')')
	}

	return ast.WalkContinue, nil
}

func (r *Renderer) renderLink(w util.BufWriter, _ []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.Link)

	if entering {
		r.write(w, '[')
	} else {
		r.write(w, "](")
		r.write(w, n.Destination)

		if n.Title != nil {
			r.write(w, " \"")
			r.write(w, n.Title)
			r.write(w, "\"")
		}
		r.write(w, ')')
	}

	return ast.WalkContinue, nil
}

func (r *Renderer) renderRawHTML(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		return ast.WalkSkipChildren, nil
	}

	r.write(w, "<!-- raw HTML omitted -->")

	return ast.WalkSkipChildren, nil
}

func (r *Renderer) renderString(w util.BufWriter, _ []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		return ast.WalkContinue, nil
	}

	n := node.(*ast.String)

	// TODO: understand if there is any risk associated with this.
	r.write(w, html.UnescapeString(string(n.Value)))

	return ast.WalkContinue, nil
}

func (r *Renderer) renderText(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		return ast.WalkContinue, nil
	}

	n := node.(*ast.Text)
	segment := n.Segment
	value := segment.Value(source)

	r.write(w, value)
	if n.HardLineBreak() {
		r.write(w, "\n\n")
	} else if n.SoftLineBreak() {
		r.write(w, '\n')
	}

	return ast.WalkContinue, nil
}
