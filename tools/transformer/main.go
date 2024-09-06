//go:generate go run ./hack/generate-directives directives.go
package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"

	"github.com/grafana/killercoda/tools/transformer/killercoda"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/text"
)

const command = "transformer"

func usage(w io.Writer, fs *flag.FlagSet) {
	fmt.Fprintf(w, "Usage of %s:\n", command)
	fs.PrintDefaults()

	fmt.Fprintln(w, "  <SOURCE FILE PATH>")
	fmt.Fprintln(w, "    	Path to the documentation source file")

	fmt.Fprintln(w, "  <DESTINATION DIRECTORY PATH>")
	fmt.Fprintln(w, "    	Path to the Killercoda output directory")
}

func main() {
	const (
		requiredSrcFilePath = iota
		requiredDstDirPath
		requiredTotal
	)

	fs := flag.NewFlagSet(command, flag.ExitOnError)
	flag.Parse()

	if flag.NArg() != requiredTotal {
		usage(os.Stderr, fs)

		os.Exit(2)
	}

	srcFilePath := flag.Arg(requiredSrcFilePath)
	dstDirPath := flag.Arg(requiredDstDirPath)

	if err := transform(srcFilePath, dstDirPath); err != nil {
		if e, ok := err.(interface{ Unwrap() []error }); ok {
			for _, err := range e.Unwrap() {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			}
		} else {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		}

		os.Exit(1)
	}
}

func writeFile(md goldmark.Markdown, dstDirPath, filename string, data []byte) error {
	root := md.Parser().Parse(text.NewReader(data))

	out, err := os.Create(filepath.Join(dstDirPath, filename))
	if err != nil {
		return fmt.Errorf("couldn't create intro file: %w", err)
	}

	if err := md.Renderer().Render(out, data, root); err != nil {
		return fmt.Errorf("couldn't render intro output: %w", err)
	}

	return nil
}

func transform(srcFilePath, dstDirPath string) error {
	if err := os.MkdirAll(dstDirPath, os.ModePerm); err != nil {
		return fmt.Errorf("couldn't create output directory: %w", err)
	}

	data, err := os.ReadFile(srcFilePath)
	if err != nil {
		return fmt.Errorf("couldn't open source file: %w", err)
	}

	md := goldmark.New(goldmark.WithExtensions(&KillercodaExtension{
		Transformers:        DefaultKillercodaTransformers,
		AdditionalExtenders: []goldmark.Extender{},
	}))

	root := md.Parser().Parse(text.NewReader(data))

	meta, ok := root.OwnerDocument().Meta()["killercoda"].(map[any]any)
	if !ok {
		return fmt.Errorf("couldn't find metadata in source file front matter")
	}

	sp, err := NewSubstitutionPreprocessorFromMeta(meta)
	if err != nil {
		return fmt.Errorf("couldn't create substitution preprocessor: %w", err)
	}

	pp := NewComposedPreprocessor(sp, NewAdmonitionPreprocessor(), NewDocsIgnorePreprocessor())

	data, err = pp.Process(data)
	if err != nil {
		return fmt.Errorf("couldn't process substitutions: %w", err)
	}

	var (
		wroteIntro  bool
		wroteFinish bool
	)

	if bytes.Contains(data, []byte(pageIntroStartMarker)) {
		md := goldmark.New(goldmark.WithExtensions(&KillercodaExtension{
			Transformers: DefaultKillercodaTransformers,
			AdditionalExtenders: []goldmark.Extender{
				&StepTransformer{StartMarker: pageIntroStartMarker, EndMarker: pageIntroEndMarker},
			},
		}))

		if err := writeFile(md, dstDirPath, "intro.md", data); err != nil {
			return err
		}

		wroteIntro = true
	}

	if bytes.Contains(data, []byte(pageFinishStartMarker)) {
		md := goldmark.New(goldmark.WithExtensions(&KillercodaExtension{
			Transformers: DefaultKillercodaTransformers,
			AdditionalExtenders: []goldmark.Extender{
				&StepTransformer{StartMarker: pageFinishStartMarker, EndMarker: pageFinishEndMarker},
			},
		}))

		if err := writeFile(md, dstDirPath, "finish.md", data); err != nil {
			return err
		}

		wroteFinish = true
	}

	var (
		errs  error
		steps int
	)

	for i := 1; i <= 20; i++ {
		startMarker := pageStepStartMarkers[i-1]
		endMarker := pageStepEndMarkers[i-1]

		if regexp.MustCompile(startMarker).Match(data) {
			steps++
			md := goldmark.New(goldmark.WithExtensions(&KillercodaExtension{
				Transformers: DefaultKillercodaTransformers,
				AdditionalExtenders: []goldmark.Extender{
					&StepTransformer{StartMarker: startMarker, EndMarker: endMarker},
				},
			}))

			if err := writeFile(md, dstDirPath, fmt.Sprintf("step%d.md", i), data); err != nil {
				errs = errors.Join(errs, err)
			}

			continue
		}

		break
	}

	if err := writeIndex(dstDirPath, meta, steps, wroteIntro, wroteFinish); err != nil {
		return err
	}

	return errs
}

func writeIndex(dstDirPath string, meta map[any]any, steps int, wroteIntro bool, wroteFinish bool) error {
	index, err := killercoda.FromMeta(meta)
	if err != nil {
		return fmt.Errorf("couldn't parse metadata: %w", err)
	}

	if wroteIntro {
		index.Details.Intro.Text = "intro.md"
	}

	for i := 0; i < steps; i++ {
		index.Details.Steps = append(index.Details.Steps, killercoda.Text{
			Text: fmt.Sprintf("step%d.md", i+1),
		})
	}

	if wroteFinish {
		index.Details.Finish.Text = "finish.md"
	}

	f, err := os.Create(filepath.Join(dstDirPath, "index.json"))
	if err != nil {
		return fmt.Errorf("couldn't create index file: %w", err)
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")

	if err := enc.Encode(index); err != nil {
		return fmt.Errorf("couldn't encode index file: %w", err)
	}

	return nil
}
