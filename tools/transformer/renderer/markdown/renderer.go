// Package renderer implements Goldmark renderer that outputs Markdown.
// This package borrows code from https://github.com/yuin/goldmark/blob/v1.7.2/renderer/html/html.go
// All borrowed code is in goldmark.go.
package markdown

import (
	"fmt"
	"strings"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

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

// TODO: replace with implementation of renderer.Writer interface.
func (r *Renderer) Write(w util.BufWriter, writee any) {
	fmt.Printf("Writing %q with indent %d, previously wrote %q\n", writee, r.indent, r.lastWrittenByte)
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
		_, _ = w.WriteString(writee)
		r.lastWrittenByte = writee[len(writee)-1]
	case []uint8:
		_, _ = w.Write(writee)
		r.lastWrittenByte = writee[len(writee)-1]
	default:
		panic(fmt.Sprintf("Write: unsupported type %T", writee))
	}
}

func (r *Renderer) writeLines(w util.BufWriter, source []byte, n ast.Node) {
	for i := 0; i < n.Lines().Len(); i++ {
		line := n.Lines().At(i)
		r.Write(w, line.Value(source))
	}
}

type Config struct {
	Attributes bool
	Unsafe     bool
	Writer     Writer
}

// NewConfig returns a new Config with defaults.
func NewConfig() Config {
	return Config{
		Attributes: false,
		Unsafe:     false,
		Writer:     DefaultWriter,
	}
}

// SetOption implements renderer.NodeRenderer.SetOption.
func (c *Config) SetOption(name renderer.OptionName, value interface{}) {
	switch name {
	case optTextWriter:
		c.Writer = value.(Writer)
	case optUnsafe:
		c.Unsafe = value.(bool)
	}
}

// An Option interface sets options for Markdown based renderers.
type Option interface {
	SetMarkdownOption(c *Config)
}

type Renderer struct {
	Config

	indent          int
	lastWrittenByte byte
}

// TextWriter is the option name for WithWriter.
const optTextWriter renderer.OptionName = "Writer"

type withWriter struct {
	value Writer
}

func (o *withWriter) SetConfig(c *renderer.Config) {
	c.Options[optTextWriter] = o.value
}

func (o *withWriter) SetMarkdownOption(c *Config) {
	c.Writer = o.value
}

// WithWriter sets the renderer's output writer.
func WithWriter(writer Writer) interface {
	renderer.Option
	Option
} {
	return &withWriter{writer}
}

// Unsafe is the option name for WithUnsafe.
const optUnsafe renderer.OptionName = "Unsafe"

type withUnsafe struct{}

func (o *withUnsafe) SetConfig(c *renderer.Config) {
	c.Options[optUnsafe] = true
}

func (o *withUnsafe) SetMarkdownOption(c *Config) {
	c.Unsafe = true
}

// WithUnsafe renders dangerous content like raw HTML and links as it is.
func WithUnsafe() interface {
	renderer.Option
	Option
} {
	return &withUnsafe{}
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
