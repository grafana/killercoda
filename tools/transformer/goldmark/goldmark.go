package goldmark

import (
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"mvdan.cc/xurls/v2"
)

// NewMarkdown returns a Goldmark Markdown interface configured to match the Grafana website settings.
// For list of default extension: https://gohugo.io/getting-started/configuration-markup/.
// For website configuration:
// https://github.com/grafana/website/blob/master/config/_default/config.yaml#L103-L121
func NewMarkdown() goldmark.Markdown {
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
	)
}
