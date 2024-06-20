// Package renderer implements Goldmark renderer that outputs Markdown.
package markdown

import (
	"fmt"
	"strings"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

// RegisterFuncs implements NodeRenderer.RegisterFuncs.
func (r *Renderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	// Block elements.
	reg.Register(ast.KindDocument, r.renderDocument)
	reg.Register(ast.KindHeading, r.renderHeading)
	reg.Register(ast.KindBlockquote, r.renderBlockquote)
	reg.Register(ast.KindCodeBlock, r.renderCodeBlock)
	reg.Register(ast.KindFencedCodeBlock, r.renderFencedCodeBlock)
	reg.Register(ast.KindHTMLBlock, r.renderHTMLBlock)
	reg.Register(ast.KindList, r.renderList)
	reg.Register(ast.KindListItem, r.renderListItem)
	reg.Register(ast.KindParagraph, r.renderParagraph)
	reg.Register(ast.KindTextBlock, r.renderTextBlock)
	reg.Register(ast.KindThematicBreak, r.renderThematicBreak)

	// Inline elements.
	reg.Register(ast.KindAutoLink, r.renderAutoLink)
	reg.Register(ast.KindCodeSpan, r.renderCodeSpan)
	reg.Register(ast.KindEmphasis, r.renderEmphasis)
	reg.Register(ast.KindImage, r.renderImage)
	reg.Register(ast.KindLink, r.renderLink)
	reg.Register(ast.KindRawHTML, r.renderRawHTML)
	reg.Register(ast.KindText, r.renderText)
	reg.Register(ast.KindString, r.renderString)
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

// write writes the current Markdown indentation prefix before dispatching writes for supported writees to the buf writer.
func (r *Renderer) write(w util.BufWriter, writee any) {
	if r.lastWrittenByte == '\n' && !isNewline(writee) {
		_, _ = w.WriteString(strings.Repeat(" ", r.indent))
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

func (r *Renderer) writeLines(w util.BufWriter, source []byte, n ast.Node) {
	for i := 0; i < n.Lines().Len(); i++ {
		line := n.Lines().At(i)
		r.write(w, line.Value(source))
	}
}

type Config struct {
	KillercodaActions bool
}

// NewConfig returns a new Config with defaults.
func NewConfig() Config {
	return Config{
		KillercodaActions: false,
	}
}

// SetOption implements renderer.NodeRenderer.SetOption.
func (c *Config) SetOption(name renderer.OptionName, value interface{}) {
	switch name {
	case optKillercodaActions:
		c.KillercodaActions = value.(bool)
	}
}

// An Option interface sets options for Markdown based renderers.
type Option interface {
	SetMarkdownOption(c *Config)
}

const optKillercodaActions renderer.OptionName = "KillercodaActions"

type withKillercodaActions struct {
}

func (o *withKillercodaActions) SetConfig(c *renderer.Config) {
	c.Options[optKillercodaActions] = true
}

func (o *withKillercodaActions) SetMarkdownOption(c *Config) {
	c.KillercodaActions = true
}

// WithKillercodaActions decides whether to render Killercoda actions for fenced code blocks.
// Actions include {{exec}} and {{copy}}.
func WithKillercodaActions() interface {
	renderer.Option
	Option
} {
	return &withKillercodaActions{}
}

type Renderer struct {
	Config

	indent          int
	lastWrittenByte byte
}

// NewRenderer configures a new Goldmark renderer for Markdown.
func NewRenderer(opts ...Option) renderer.NodeRenderer {
	renderer := &Renderer{
		Config: NewConfig(),

		indent:          0,
		lastWrittenByte: 0,
	}

	for _, opt := range opts {
		opt.SetMarkdownOption(&renderer.Config)
	}
	return renderer
}
