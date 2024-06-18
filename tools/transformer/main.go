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

	copyStartMarker    = "<!-- Killercoda copy START -->"
	copyEndMarker      = "<!-- Killercoda copy END -->"
	executeStartMarker = "<!-- Killercoda execute START -->"
	executeEndMarker   = "<!-- Killercoda execute END -->"
)

func writeIntro(data []byte) {
	md := goldmark.NewMarkdown()
	md.Parser().AddOptions(parser.WithASTTransformers(
		util.Prioritized(&IgnoreTransformer{}, 0),
		util.Prioritized(&IntroTransformer{}, 1),
	))
	md.SetRenderer(renderer.NewRenderer(renderer.WithNodeRenderers(util.Prioritized(markdown.NewRenderer(), 1000))))

	root := md.Parser().Parse(text.NewReader(data))

	outFile := filepath.Join(dstPath, "intro.md")

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

	writeIntro(data)
}
