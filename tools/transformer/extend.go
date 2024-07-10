package main

import (
	tgoldmark "github.com/grafana/killercoda/tools/transformer/goldmark"
	"github.com/grafana/killercoda/tools/transformer/goldmark/renderer/markdown"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/util"
)

var DefaultKillercodaTransformers = []util.PrioritizedValue{
	util.Prioritized(&IgnoreTransformer{}, 1),
	util.Prioritized(&FigureTransformer{}, 2),
	util.Prioritized(&InlineActionTransformer{}, 3),
	util.Prioritized(&ActionTransformer{Kind: "copy"}, 3),
	util.Prioritized(&ActionTransformer{Kind: "exec"}, 3),
	util.Prioritized(&LinkTransformer{}, 4),
	util.Prioritized(&HeadingTransformer{}, 5),
}

// KillercodaExtension extends the Goldmark Markdown parser with the transformations that convert Hugo Markdown to KillercodaExtension Markdown.
type KillercodaExtension struct {
	Transformers        []util.PrioritizedValue
	AdditionalExtenders []goldmark.Extender
}

// Extend implements the goldmark.Extender interface.
// It adds the default AST transformers that convert Hugo Markdown to Killercoda Markdown.
func (e *KillercodaExtension) Extend(md goldmark.Markdown) {
	tgoldmark.NewWebsite().Extend(md)

	md.Parser().AddOptions(
		parser.WithASTTransformers(
			e.Transformers...))

	markdown.NewRenderer(
		markdown.WithKillercodaActions(),
	).Extend(md)

	for _, extender := range e.AdditionalExtenders {
		extender.Extend(md)
	}
}
