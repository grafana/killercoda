package parser

import (
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	goldmarkHTML "github.com/yuin/goldmark/renderer/html"
	"mvdan.cc/xurls/v2"
)

// New returns a Markdown parser configured to match the website parser.
// For list of default extension: https://gohugo.io/getting-started/configuration-markup/.
// For website configuration:
// https://github.com/grafana/website/blob/master/config/_default/config.yaml#L103-L121
func New() parser.Parser {
	return goldmark.New(
		goldmark.WithExtensions(
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
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			goldmarkHTML.WithHardWraps(),
			goldmarkHTML.WithUnsafe(),
			goldmarkHTML.WithXHTML(),
		),
	).Parser()
}
