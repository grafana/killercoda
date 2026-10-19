// Package renderer implements Goldmark renderer that outputs Markdown.
package markdown

import (
	"fmt"
	"io"
	"strings"

	tableext "github.com/grafana/killercoda/tools/transformer/goldmark/extension"
	"github.com/yuin/goldmark/v2/ast"
	"github.com/yuin/goldmark/v2/renderer"
	"github.com/yuin/goldmark/v2/util"
)

// Type aliases.
type Option = renderer.Option[Config]

type Config struct {
	renderer.Config[util.BufWriter, Config]
	KillercodaActions bool
	Unsafe            bool
}

// Default returns a Config with default values.
func (c Config) Default() Config {
	return Config{
		KillercodaActions: false,
		Unsafe:            false,
	}
}

// WithKillercodaActions decides whether to render Killercoda actions for fenced code blocks.
// Actions include {{exec}} and {{copy}}.
func WithKillercodaActions() Option {
	return renderer.NewOptionFunc(func(c *Config) {
		c.KillercodaActions = true
	})
}

// WithUnsafe decides whether to render unsafe HTML.
func WithUnsafe() Option {
	return renderer.NewOptionFunc(func(c *Config) {
		c.Unsafe = true
	})
}

type Renderer struct {
	*renderer.Helper[util.BufWriter, Config]

	prefix          string
	lastWrittenByte byte
}

// NewRenderer configures a new Goldmark renderer for Markdown.
func NewRenderer(opts ...Option) *Renderer {
	r := &Renderer{}

	nodes := map[ast.NodeKind]renderer.NodeRenderer[util.BufWriter]{
		ast.KindDocument:      renderer.NodeRendererFunc(r.renderDocument),
		ast.KindHeading:       renderer.NodeRendererFunc(r.renderHeading),
		ast.KindBlockquote:    renderer.NodeRendererFunc(r.renderBlockquote),
		ast.KindCodeBlock:     renderer.NodeRendererFunc(r.renderCodeBlock),
		ast.KindHTMLBlock:     renderer.NodeRendererFunc(r.renderHTMLBlock),
		ast.KindList:          renderer.NodeRendererFunc(r.renderList),
		ast.KindListItem:      renderer.NodeRendererFunc(r.renderListItem),
		ast.KindParagraph:     renderer.NodeRendererFunc(r.renderParagraph),
		ast.KindThematicBreak: renderer.NodeRendererFunc(r.renderThematicBreak),

		ast.KindAutoLink: renderer.NodeRendererFunc(r.renderAutoLink),
		ast.KindCodeSpan: renderer.NodeRendererFunc(r.renderCodeSpan),
		ast.KindEmphasis: renderer.NodeRendererFunc(r.renderEmphasis),
		ast.KindStrong:   renderer.NodeRendererFunc(r.renderStrong),
		ast.KindImage:    renderer.NodeRendererFunc(r.renderImage),
		ast.KindLink:     renderer.NodeRendererFunc(r.renderLink),
		ast.KindRawHTML:  renderer.NodeRendererFunc(r.renderRawHTML),
		ast.KindText:     renderer.NodeRendererFunc(r.renderText),
	}

	for kind, nodeRenderer := range tableext.TableNodeRenderers() {
		nodes[kind] = nodeRenderer
	}

	opts = append([]Option{renderer.WithNodeRenderers[util.BufWriter, Config](nodes)}, opts...)
	r.Helper = renderer.NewHelper[util.BufWriter](opts...)

	return r
}

// Render renders the given AST node to the given writer.
func (r *Renderer) Render(w io.Writer, source []byte, n ast.Node, opts ...renderer.RenderOption) error {
	r.prefix = ""
	r.lastWrittenByte = 0

	if ew, ok := w.(util.ErrorBufWriter); ok {
		return r.Helper.Render(ew, source, n, opts...)
	}

	return r.Helper.Render(util.NewErrorBufWriterSize(w, len(source)*3), source, n, opts...)
}

// isNewline checks whether the writee is a single newline character.
func isNewline(writee any) bool {
	switch writee := writee.(type) {
	case byte:
		return writee == '\n'
	case int32:
		return writee == '\n'
	case string:
		return len(writee) == 1 && writee[0] == '\n'
	case []uint8:
		return len(writee) == 1 && writee[0] == '\n'
	default:
		panic(fmt.Sprintf("Write: unsupported type %T", writee))
	}
}

// write writes the current Markdown prefix before dispatching writes for supported writees to the buf writer.
func (r *Renderer) write(w util.BufWriter, writee any) {
	if r.lastWrittenByte == '\n' {
		if isNewline(writee) {
			w.WriteString(strings.TrimSpace(r.prefix))
		} else {
			w.WriteString(r.prefix)
		}
	}

	switch writee := writee.(type) {
	case byte:
		_ = w.WriteByte(writee)
		r.lastWrittenByte = writee
	case int32:
		_ = w.WriteByte(byte(writee))
		r.lastWrittenByte = byte(writee)
	case string:
		if len(writee) == 0 {
			return
		}

		_, _ = w.WriteString(writee)
		r.lastWrittenByte = writee[len(writee)-1]
	case []uint8:
		if len(writee) == 0 {
			return
		}
		_, _ = w.Write(writee)
		r.lastWrittenByte = writee[len(writee)-1]
	default:
		panic(fmt.Sprintf("Write: unsupported type %T", writee))
	}
}

// secureWrite writes the source to the buf writer, replacing any null characters with the replacement character.
func (r *Renderer) secureWrite(w util.BufWriter, source []byte) {
	var (
		n int
		l = len(source)
	)

	for i := 0; i < l; i++ {
		if source[i] == '\u0000' {
			r.write(w, source[i-n:i])
			n = 0
			r.write(w, []byte("\ufffd"))

			continue
		}

		n++
	}

	if n != 0 {
		r.write(w, source[l-n:])
	}
}

func (r *Renderer) writeLines(w util.BufWriter, source []byte, lines [][]byte) {
	for _, line := range lines {
		r.write(w, line)
	}
}

func (r *Renderer) pushPrefix(str string) {
	r.prefix += str
}

func (r *Renderer) popPrefix(str string) {
	r.prefix = strings.TrimSuffix(r.prefix, str)
}
