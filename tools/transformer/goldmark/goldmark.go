// Package goldmark provides parser configuration helpers.
package goldmark

import (
	meta "github.com/yuin/goldmark-meta/v2"
	"github.com/yuin/goldmark/v2/extension"
	"github.com/yuin/goldmark/v2/parser"
	"mvdan.cc/xurls/v2"
)

// NewWebsiteParser returns a parser configured to approximate the Grafana website configuration.
// For list of default extensions: https://gohugo.io/getting-started/configuration-markup/.
// For the website configuration:
// https://github.com/grafana/website/blob/master/config/_default/config.yaml#L103-L121
func NewWebsiteParser(opts ...parser.Option) parser.Parser {
	base := []parser.Option{
		parser.WithExtensions(
			meta.Parser,
			extension.NewDefinitionListParser(),
			extension.NewFootnoteParser(),
			extension.NewLinkifyParser(
				extension.WithAllowedProtocols([][]byte{
					[]byte("http:"),
					[]byte("https:"),
				}),
				extension.WithURLRegexp(xurls.Strict()),
			),
			extension.NewStrikethroughParser(),
			extension.NewTableParser(),
			extension.NewTaskListItemParser(),
		),
		parser.WithAutoHeadingID(),
	}

	base = append(base, opts...)

	return parser.New(base...)
}
