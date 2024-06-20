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

	"github.com/grafana/killercoda/tools/transformer/goldmark"
	"github.com/grafana/killercoda/tools/transformer/goldmark/renderer/markdown"
	"github.com/grafana/killercoda/tools/transformer/killercoda"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

const (
	command = "transformer"

	indexFilename  = "index.json"
	introFilename  = "intro.md"
	finishFilename = "finish.md"
)

func usage(w io.Writer, fs *flag.FlagSet) {
	fmt.Fprintf(w, "Usage of %s:\n", command)
	fs.PrintDefaults()

	fmt.Fprintln(w, "  <SOURCE FILE PATH>")
	fmt.Fprintln(w, "    	Path to the documentation source file")

	fmt.Fprintln(w, "  <DESTINATION DIRECTORY PATH>")
	fmt.Fprintln(w, "    	Path to the Killercoda output directory")
}

func writeIntro(dstDirPath string, data []byte, transformers []util.PrioritizedValue, renderer renderer.Renderer) error {
	md := goldmark.NewMarkdown()
	md.Parser().AddOptions(
		parser.WithASTTransformers(
			append(transformers, util.Prioritized(&IntroTransformer{}, 0))...),
	)
	md.SetRenderer(renderer)

	root := md.Parser().Parse(text.NewReader(data))

	out, err := os.Create(filepath.Join(dstDirPath, introFilename))
	if err != nil {
		return fmt.Errorf("couldn't create intro file: %w", err)
	}

	if err := md.Renderer().Render(out, data, root); err != nil {
		return fmt.Errorf("couldn't render intro output: %w", err)
	}

	return nil
}

func writeFinish(dstDirPath string, data []byte, transformers []util.PrioritizedValue, renderer renderer.Renderer) error {
	md := goldmark.NewMarkdown()
	md.Parser().AddOptions(
		parser.WithASTTransformers(
			append(transformers, util.Prioritized(&FinishTransformer{}, 0))...),
	)
	md.SetRenderer(renderer)

	root := md.Parser().Parse(text.NewReader(data))

	out, err := os.Create(filepath.Join(dstDirPath, finishFilename))
	if err != nil {
		return fmt.Errorf("couldn't create finish file: %w", err)
	}

	if err := md.Renderer().Render(out, data, root); err != nil {
		return fmt.Errorf("couldn't render finish output: %w", err)
	}

	return nil
}

func writeStep(dstDirPath string, n int, data []byte, transformers []util.PrioritizedValue, renderer renderer.Renderer) error {
	md := goldmark.NewMarkdown()
	md.Parser().AddOptions(
		parser.WithASTTransformers(
			append(transformers, util.Prioritized(&StepTransformer{Step: n}, 0))...),
	)
	md.SetRenderer(renderer)

	root := md.Parser().Parse(text.NewReader(data))

	out, err := os.Create(filepath.Join(dstDirPath, fmt.Sprintf("step%d.md", n)))
	if err != nil {
		return fmt.Errorf("could create step file: %w", err)
	}

	if err := md.Renderer().Render(out, data, root); err != nil {
		return fmt.Errorf("couldn't render step %d output: %w", n, err)
	}

	return nil
}

func writeIndex(dstDirPath string, meta map[any]any, steps int, wroteIntro bool, wroteFinish bool) error {
	index, err := killercoda.FromMeta(meta)
	if err != nil {
		return fmt.Errorf("couldn't parse metadata: %w", err)
	}

	if wroteIntro {
		index.Details.Intro.Text = introFilename
	}

	for i := 0; i < steps; i++ {
		index.Details.Steps = append(index.Details.Steps, killercoda.Text{
			Text: fmt.Sprintf("step%d.md", i+1),
		})
	}

	if wroteFinish {
		index.Details.Finish.Text = finishFilename
	}

	f, err := os.Create(filepath.Join(dstDirPath, indexFilename))
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

func transform(srcFilePath, dstDirPath string) error {
	data, err := os.ReadFile(srcFilePath)
	if err != nil {
		return fmt.Errorf("couldn't open source file: %w", err)
	}

	md := goldmark.NewMarkdown()
	root := md.Parser().Parse(text.NewReader(data))

	meta, ok := root.OwnerDocument().Meta()["killercoda"].(map[any]any)
	if !ok {
		return fmt.Errorf("couldn't find metadata in source file front matter")
	}

	if preprocessing, ok := meta["preprocessing"].(map[any]any); ok {
		if substitutions, ok := preprocessing["substitutions"].([]any); ok {
			for _, substitution := range substitutions {
				if s, ok := substitution.(map[any]any); ok {
					if expr, ok := s["regexp"].(string); ok {
						if replacement, ok := s["replacement"].(string); ok {
							data = regexp.MustCompile(expr).ReplaceAll(data, []byte(replacement))
						}
					}
				}
			}
		}
	}

	transformers := []util.PrioritizedValue{
		util.Prioritized(&IgnoreTransformer{}, 1),
		util.Prioritized(&IncludeTransformer{}, 1),
		util.Prioritized(&FigureTransformer{}, 2),
		util.Prioritized(&InlineActionTransformer{}, 3),
		util.Prioritized(&ActionTransformer{Kind: "copy"}, 3),
		util.Prioritized(&ActionTransformer{Kind: "exec"}, 3),
		util.Prioritized(&LinkTransformer{}, 4),
		util.Prioritized(&HeadingTransformer{}, 5),
	}

	//nolint:gomnd // These priority values are relative to each other and are not magic.
	renderer := renderer.NewRenderer(
		renderer.WithNodeRenderers(
			util.Prioritized(
				markdown.NewRenderer(
					markdown.WithKillercodaActions(),
				), 1000)))

	var (
		wroteIntro  bool
		wroteFinish bool
	)

	if bytes.Contains(data, []byte(introStartMarker)) {
		if err := writeIntro(dstDirPath, data, transformers, renderer); err != nil {
			return err
		}

		wroteIntro = true
	}

	if bytes.Contains(data, []byte(finishStartMarker)) {
		if err := writeFinish(dstDirPath, data, transformers, renderer); err != nil {
			return err
		}

		wroteFinish = true
	}

	var (
		errs  error
		steps int
	)

	for i := 1; i <= 5; i++ {
		if regexp.MustCompile(fmt.Sprintf(`<!-- Killercoda step%d.md START -->`, i)).Match(data) {
			steps++

			if err := writeStep(dstDirPath, i, data, transformers, renderer); err != nil {
				errs = errors.Join(errs, err)
			}
		}
	}

	writeIndex(dstDirPath, meta, steps, wroteIntro, wroteFinish)

	return errs
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
