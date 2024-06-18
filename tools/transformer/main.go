package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/grafana/killercoda/tools/transformer/parser"
	"github.com/grafana/killercoda/tools/transformer/renderer/markdown"
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
	finishStartMarker  = "<!-- Killercoda finish.md START -->"
	finishEndMarker    = "<!-- Killercoda finish.md END -->"
	ignoreStartMarker  = "<!-- Killercoda ignore START -->"
	ignoreEndMarker    = "<!-- Killercoda ignore END -->"
	introStartMarker   = "<!-- Killercoda intro.md START -->"
	introEndMarker     = "<!-- Killercoda intro.md END -->"
)

func main() {
	data, err := os.ReadFile(srcPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't open source file: %v\n", err)

		os.Exit(1)
	}

	root := parser.New().Parse(text.NewReader(data))

	outFile := filepath.Join(dstPath, "step1.md")
	out, err := os.Create(outFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't create output file: %v\n", err)
	}

	r := renderer.NewRenderer(renderer.WithNodeRenderers(util.Prioritized(markdown.NewRenderer(), 1000)))
	r.Render(out, data, root)
}
