// Package goldmark provides extensions to the Goldmark Markdown interface.
package goldmark

import (
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"mvdan.cc/xurls/v2"
)

// Extension contains Goldmark extensions.
type Extension struct {
	Extenders     []goldmark.Extender
	ParserOptions []parser.Option
}

// NewWebsite returns an extension to the Goldmark Markdown interface configured to approximate the Grafana website configuration.
// For list of default extensions: https://gohugo.io/getting-started/configuration-markup/.
// For the website configuration:
// https://github.com/grafana/website/blob/master/config/_default/config.yaml#L103-L121
func NewWebsite() *Extension {
	return &Extension{
		Extenders: []goldmark.Extender{
			extension.DefinitionList,
			extension.Footnote,
			extension.NewLinkify(
				extension.WithLinkifyAllowedProtocols([][]byte{
					[]byte("http:"),
					[]byte("https:"),
				}),
				extension.WithLinkifyURLRegexp(
					xurls.Strict(),
				)),
			extension.Strikethrough,
			extension.Table,
			extension.TaskList,
			extension.Typographer,
			meta.New(meta.WithStoresInDocument()),
		},
		ParserOptions: []parser.Option{
			parser.WithAutoHeadingID(),
		},
	}
}

func (e *Extension) Extend(md goldmark.Markdown) {
	md.Parser().AddOptions(e.ParserOptions...)

	for _, extender := range e.Extenders {
		extender.Extend(md)
	}
}
