package main

import (
	"io"

	tgoldmark "github.com/grafana/killercoda/tools/transformer/goldmark"
	gk "github.com/grafana/killercoda/tools/transformer/goldmark/killercoda"
	"github.com/grafana/killercoda/tools/transformer/goldmark/renderer/markdown"
	"github.com/yuin/goldmark/v2/ast"
	"github.com/yuin/goldmark/v2/parser"
	"github.com/yuin/goldmark/v2/util"
)

var DefaultKillercodaTransformers = []util.PrioritizedValue[parser.ASTTransformer]{
	util.Prioritized[parser.ASTTransformer](&IgnoreTransformer{}, 1),
	util.Prioritized[parser.ASTTransformer](&FigureTransformer{}, 2),
	util.Prioritized[parser.ASTTransformer](&InlineActionTransformer{}, 3),
	util.Prioritized[parser.ASTTransformer](&ActionTransformer{Kind: "copy"}, 3),
	util.Prioritized[parser.ASTTransformer](&ActionTransformer{Kind: "exec"}, 3),
	util.Prioritized[parser.ASTTransformer](&LinkTransformer{}, 4),
	util.Prioritized[parser.ASTTransformer](&HeadingTransformer{}, 5),
}

type TransformerMarkdown struct {
	parser   parser.Parser
	renderer *markdown.Renderer
}

func NewTransformerMarkdown(base []util.PrioritizedValue[parser.ASTTransformer], additional ...parser.ASTTransformer) *TransformerMarkdown {
	transformers := append([]util.PrioritizedValue[parser.ASTTransformer]{}, base...)
	for i, transformer := range additional {
		transformers = append(transformers, util.Prioritized[parser.ASTTransformer](transformer, i))
	}

	p := tgoldmark.NewWebsiteParser(
		parser.WithBlockParsers(
			util.Prioritized[parser.BlockParser](gk.NewFencedCodeBlockParser(), 101),
		),
		parser.WithASTTransformers(transformers...),
	)

	return &TransformerMarkdown{
		parser:   p,
		renderer: markdown.NewRenderer(markdown.WithKillercodaActions()),
	}
}

func (m *TransformerMarkdown) Parse(source []byte) ast.Node {
	return m.parser.Parse(source)
}

func (m *TransformerMarkdown) Render(w io.Writer, source []byte, n ast.Node) error {
	return m.renderer.Render(w, source, n)
}
