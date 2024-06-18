package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/grafana/killercoda/tools/transformer/goldmark"
	"github.com/grafana/killercoda/tools/transformer/goldmark/renderer/markdown"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

const (
	srcPath = "/Users/jdb/ext/grafana/loki/docs/sources/get-started/quick-start.md"
	dstPath = "/Users/jdb/ext/grafana/killercoda/loki/loki-quickstart"

	finishFile = "finished.md"
	indexFile  = "index.json"
	introFile  = "intro.md"
)

func writeIntro(data []byte, renderer renderer.Renderer) {
	md := goldmark.NewMarkdown()
	md.Parser().AddOptions(parser.WithASTTransformers(
		util.Prioritized(&IntroTransformer{}, 0),
		util.Prioritized(&IgnoreTransformer{}, 1),
		util.Prioritized(&LinkTransformer{}, 2),
		util.Prioritized(&CopyTransformer{}, 3),
		util.Prioritized(&ExecTransformer{}, 3),
	))
	md.SetRenderer(renderer)

	root := md.Parser().Parse(text.NewReader(data))

	outFile := filepath.Join(dstPath, "intro.md")

	out, err := os.Create(outFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't create output file: %v\n", err)
	}

	md.Renderer().Render(out, data, root)
}

func writeStepOne(data []byte, renderer renderer.Renderer) {
	md := goldmark.NewMarkdown()
	md.Parser().AddOptions(parser.WithASTTransformers(
		util.Prioritized(&StepTransformer{}, 0),
		util.Prioritized(&IgnoreTransformer{}, 1),
		util.Prioritized(&LinkTransformer{}, 2),
		util.Prioritized(&CopyTransformer{}, 3),
		util.Prioritized(&ExecTransformer{}, 3),
	))
	md.SetRenderer(renderer)

	root := md.Parser().Parse(text.NewReader(data))

	outFile := filepath.Join(dstPath, "step1.md")

	out, err := os.Create(outFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't create output file: %v\n", err)
	}

	md.Renderer().Render(out, data, root)
}

func main() {
	data, err := os.ReadFile(srcPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't open source file: %v\n", err)

		os.Exit(1)
	}

	renderer := renderer.NewRenderer(
		renderer.WithNodeRenderers(
			util.Prioritized(
				markdown.NewRenderer(
					markdown.WithKillercodaActions(),
				), 1000)))

	writeIntro(data, renderer)
	writeStepOne(data, renderer)
}
