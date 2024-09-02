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

	pp, err := NewSubstitutionPreprocessorFromMeta(meta)
	if err != nil {
		return fmt.Errorf("couldn't create substitution preprocessor: %w", err)
	}

	pp.AddSubstitution(docsIgnoreRegexp, []byte(""))

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

var docsIgnoreRegexp = regexp.MustCompile("{{< *?/?docs/ignore *?>}}\n?")

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

	// Check if forground scripts have been added to the intro or finish within meta
	// If so, copy the scripts from `sandbox-scripts` to the destination directory
	if introScript := index.Details.Intro.Foreground; introScript != "" {
		if err := copyScript(introScript, dstDirPath); err != nil {
			return fmt.Errorf("couldn't copy intro script: %w", err)
		}
	}

	if finishScript := index.Details.Finish.Foreground; finishScript != "" {
		if err := copyScript(finishScript, dstDirPath); err != nil {
			return fmt.Errorf("couldn't copy finish script: %w", err)
		}
	}

	if finishScript := index.Details.Finish.Background; finishScript != "" {
		if err := copyScript(finishScript, dstDirPath); err != nil {
			return fmt.Errorf("couldn't copy finish script: %w", err)
		}
	}

	if introScript := index.Details.Intro.Background; introScript != "" {
		if err := copyScript(introScript, dstDirPath); err != nil {
			return fmt.Errorf("couldn't copy intro script: %w", err)
		}
	}

	// Clean up any scripts which are not part of the index.json
	if err := cleanUpScripts(dstDirPath, index); err != nil {
		return fmt.Errorf("couldn't clean up scripts: %w", err)
	}

	return nil
}

// Helper function to copy a script from `sandbox-scripts` to the destination directory
func copyScript(scriptName, dstDirPath string) error {
	repoPath := "../../sandbox-scripts"
	sourcePath := filepath.Join(repoPath, scriptName)
	destPath := filepath.Join(dstDirPath, scriptName)

	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("couldn't open source script file: %w", err)
	}

	defer sourceFile.Close()

	destFile, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("couldn't create destination script file: %w", err)
	}
	defer destFile.Close()

	if _, err := io.Copy(destFile, sourceFile); err != nil {
		return fmt.Errorf("couldn't copy script file: %w", err)
	}

	return nil
}

// Helper function to clean up any scripts which are not part of the index.json
func cleanUpScripts(dstDirPath string, index killercoda.Index) error {

	scriptSet := make(map[string]struct{})

	addScript := func(script string) {
		if script != "" {
			scriptSet[script] = struct{}{}
		}
	}

	addScript(index.Details.Intro.Foreground)
	addScript(index.Details.Intro.Background)
	addScript(index.Details.Finish.Foreground)
	addScript(index.Details.Finish.Background)

	// Walk through the destination directory and remove any .sh scripts which are not part of the index
	err := filepath.Walk(dstDirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(path) == ".sh" {
			relPath, err := filepath.Rel(dstDirPath, path)
			if err != nil {
				return err
			}

			if _, exists := scriptSet[relPath]; !exists {
				if err := os.Remove(path); err != nil {
					return fmt.Errorf("couldn't remove script file %s: %w", path, err)
				}
			}
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("error cleaning up scripts: %w", err)
	}

	return nil
}
