package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"

	"github.com/grafana/killercoda/tools/transformer/goldmark"
	"github.com/grafana/killercoda/tools/transformer/goldmark/renderer/markdown"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

const (
	command = "transformer"
)

var (
	errCreateFile   = fmt.Errorf("couldn't create output file")
	errRenderOutput = fmt.Errorf("couldn't create render output")
)

func usage(w io.Writer, fs *flag.FlagSet) {
	fmt.Fprintf(w, "Usage of %s:\n", command)
	fs.PrintDefaults()

	fmt.Fprintln(w, "  <SOURCE FILE PATH>")
	fmt.Fprintln(w, "    	Path to the documentation source file")

	fmt.Fprintln(w, "  <DESTINATION DIRECTORY PATH>")
	fmt.Fprintln(w, "    	Path to the Killercoda output directory")
}

func writeIntro(dstDirPath string, data []byte, renderer renderer.Renderer) error {
	md := goldmark.NewMarkdown()
	//nolint:gomnd // These priority values are relative to each other and are not magic.
	md.Parser().AddOptions(parser.WithASTTransformers(
		util.Prioritized(&IntroTransformer{}, 0),
		util.Prioritized(&IgnoreTransformer{}, 1),
		util.Prioritized(&IncludeTransformer{}, 1),
		util.Prioritized(&FigureTransformer{}, 2),
		util.Prioritized(&InlineActionTransformer{}, 3),
		util.Prioritized(&ActionTransformer{Kind: "copy"}, 3),
		util.Prioritized(&ActionTransformer{Kind: "exec"}, 3),
		util.Prioritized(&LinkTransformer{}, 4),
		util.Prioritized(&HeadingTransformer{}, 5),
	))
	md.SetRenderer(renderer)

	root := md.Parser().Parse(text.NewReader(data))

	outFile := filepath.Join(dstDirPath, "intro.md")

	out, err := os.Create(outFile)
	if err != nil {
		return fmt.Errorf("%v: %w", errCreateFile, err)
	}

	if err := md.Renderer().Render(out, data, root); err != nil {
		return fmt.Errorf("%v: %w", errRenderOutput, err)
	}

	return nil
}

func writeStep(dstDirPath string, n int, data []byte, renderer renderer.Renderer) error {
	md := goldmark.NewMarkdown()
	//nolint:gomnd // These priority values are relative to each other and are not magic.
	md.Parser().AddOptions(parser.WithASTTransformers(
		util.Prioritized(&StepTransformer{Step: n}, 0),
		util.Prioritized(&IgnoreTransformer{}, 1),
		util.Prioritized(&IncludeTransformer{}, 1),
		util.Prioritized(&FigureTransformer{}, 2),
		util.Prioritized(&InlineActionTransformer{}, 3),
		util.Prioritized(&ActionTransformer{Kind: "copy"}, 3),
		util.Prioritized(&ActionTransformer{Kind: "exec"}, 3),
		util.Prioritized(&LinkTransformer{}, 4),
		util.Prioritized(&HeadingTransformer{}, 5),
	))
	md.SetRenderer(renderer)

	root := md.Parser().Parse(text.NewReader(data))

	outFile := filepath.Join(dstDirPath, fmt.Sprintf("step%d.md", n))

	out, err := os.Create(outFile)
	if err != nil {
		return fmt.Errorf("%v: %w", errCreateFile, err)
	}

	if err := md.Renderer().Render(out, data, root); err != nil {
		return fmt.Errorf("%v: %w", errRenderOutput, err)
	}

	return nil
}

func transform(srcFilePath, dstDirPath string) error {
	data, err := os.ReadFile(srcFilePath)
	if err != nil {
		return fmt.Errorf("couldn't open source file: %w", err)
	}

	//nolint:gomnd // These priority values are relative to each other and are not magic.
	renderer := renderer.NewRenderer(
		renderer.WithNodeRenderers(
			util.Prioritized(
				markdown.NewRenderer(
					markdown.WithKillercodaActions(),
				), 1000)))

	if err := writeIntro(dstDirPath, data, renderer); err != nil {
		return err
	}

	var errs error

	for i := 1; i <= 5; i++ {
		if regexp.MustCompile(fmt.Sprintf(`<!-- Killercoda step%d.md START -->`, i)).Match(data) {
			if err := writeStep(dstDirPath, i, data, renderer); err != nil {
				errs = errors.Join(errs, err)
			}
		}
	}

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
