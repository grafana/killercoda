package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/grafana/killercoda/tools/transformer/extract"
	"github.com/grafana/killercoda/tools/transformer/killercoda"
	"github.com/grafana/killercoda/tools/transformer/parser"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
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

func writeFinish(data []byte) {
	f, err := os.Create(filepath.Join(dstPath, finishFile))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't open finish file: %v\n", err)
	}
	defer f.Close()

	if _, err := f.Write(data); err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't write finish file: %v\n", err)
	}
}

func writeIndex(data []byte, steps int, hasIntro bool, hasFinish bool) {
	root := parser.New().Parse(text.NewReader(data))

	meta, ok := root.OwnerDocument().Meta()["killercoda"].(map[any]any)
	if !ok {
		fmt.Fprintf(os.Stderr, "No metadata found in source file\n")
	}

	index, err := killercoda.FromMeta(meta)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't parse metadata: %v\n", err)
	}

	if hasIntro {
		index.Details.Intro.Text = introFile
	}

	for i := 0; i < steps; i++ {
		index.Details.Steps = append(index.Details.Steps, killercoda.Text{
			Text: fmt.Sprintf("step%d.md", i+1),
		})
	}

	if hasFinish {
		index.Details.Finish.Text = introFile
	}

	f, err := os.Create(filepath.Join(dstPath, indexFile))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't open index file: %v\n", err)
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")

	if err := enc.Encode(index); err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't encode index file: %v\n", err)
	}
}

func writeIntro(data []byte) {
	f, err := os.Create(filepath.Join(dstPath, introFile))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't open intro file: %v\n", err)
	}
	defer f.Close()

	if _, err := f.Write(data); err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't write intro file: %v\n", err)
	}
}

func writeStep(data []byte, step int) {
	f, err := os.Create(filepath.Join(dstPath, fmt.Sprintf("step%d.md", step)))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't open step file: %v\n", err)
	}
	defer f.Close()

	if _, err := f.Write(data); err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't write step file: %v\n", err)
	}
}

func main() {
	data, err := os.ReadFile(srcPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't open source file: %v\n", err)

		os.Exit(1)
	}

	data = regexp.MustCompile(fmt.Sprintf(`(?s)%s.*?%s`, ignoreStartMarker, ignoreEndMarker)).ReplaceAll(data, []byte{})
	data = regexp.MustCompile(fmt.Sprintf(`(?s)\s*%s.*?%s`, executeStartMarker, executeEndMarker)).ReplaceAllFunc(data, func(match []byte) []byte {
		var closing bool

		match = regexp.MustCompile(fmt.Sprintf(`(?s)\s*%s(.*?)\s*%s`, executeStartMarker, executeEndMarker)).ReplaceAll(match, []byte("$1"))

		return regexp.MustCompile("```").ReplaceAllFunc(match, func(match []byte) []byte {
			if closing {
				return []byte("```{{execute}}")
			}

			closing = !closing

			return match
		})
	})
	data = regexp.MustCompile(fmt.Sprintf(`(?s)\s*%s.*?%s`, copyStartMarker, copyEndMarker)).ReplaceAllFunc(data, func(match []byte) []byte {
		var closing bool

		match = regexp.MustCompile(fmt.Sprintf(`(?s)\s*%s(.*?)\s*%s`, copyStartMarker, copyEndMarker)).ReplaceAll(match, []byte("$1"))

		return regexp.MustCompile("```").ReplaceAllFunc(match, func(match []byte) []byte {
			if closing {
				return []byte("```{{copy}}")
			}

			closing = !closing

			return match
		})
	})

	root := parser.New().Parse(text.NewReader(data))

	_ = ast.Walk(root, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		// The callback function is called twice for every node: once with entering=true when the branch is
		// first visited, then with entering=false after all the children are done. We just need to validate
		// each node once.
		if !entering {
			return ast.WalkContinue, nil
		}

		for c := node.FirstChild(); c != nil; c = c.NextSibling() {
			if c, ok := c.(*ast.Link); ok {
				fmt.Printf("Found link: %s\n", c.Destination)

				replacement := ast.NewLink()
				replacement.Destination = []byte("BLAH")
				replacement.Title = c.Title

				for gc := c.FirstChild(); gc != nil; gc = gc.NextSibling() {
					replacement.AppendChild(replacement, gc)
				}

				c.Dump(data, 0)
				replacement.Dump(data, 0)

				node.ReplaceChild(node, c, replacement)
			}
		}

		return ast.WalkContinue, nil
	})

	var (
		hasIntro  bool
		hasFinish bool
		steps     int
	)

	if bytes.Contains(data, []byte(introStartMarker)) {
		data := extract.BetweenMarkers(data, []byte(introStartMarker), []byte(introEndMarker))
		hasIntro = true

		writeIntro(data)
	}

	// Maximum of 10 steps for now.
	for i := 0; i < 10; i++ {
		stepStartMarker := fmt.Sprintf("<!-- Killercoda step%d.md START -->", i+1)
		stepEndMarker := fmt.Sprintf("<!-- Killercoda step%d.md END -->", i+1)

		if !bytes.Contains(data, []byte(stepStartMarker)) {
			break
		}

		data := extract.BetweenMarkers(data, []byte(stepStartMarker), []byte(stepEndMarker))

		writeStep(data, i+1)

		steps++
	}

	if bytes.Contains(data, []byte(finishStartMarker)) {
		data := extract.BetweenMarkers(data, []byte(finishStartMarker), []byte(finishEndMarker))
		hasFinish = true

		writeFinish(data)
	}

	writeIndex(data, steps, hasIntro, hasFinish)
}
