package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"io"
	"os"
	"strings"
)

const command = "generate-directives"

func usage(w io.Writer, fs *flag.FlagSet) {
	fmt.Fprintf(w, "Usage of %s:\n", command)
	fs.PrintDefaults()

	fmt.Fprintln(w, "  <OUTPUT PATH>")
	fmt.Fprintln(w, "    	Path to output file")
}

func main() {
	type requiredArg int

	const (
		requiredOutputPath requiredArg = iota
		requiredArgCount
	)

	flag.Parse()

	if flag.NArg() != int(requiredArgCount) {
		usage(os.Stderr, flag.CommandLine)

		os.Exit(2)
	}

	outputPath := flag.Arg(int(requiredOutputPath))

	if err := generateMarkers(outputPath); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
}

func generateMarkers(outputPath string) error {
	g := Generator{
		buf: bytes.Buffer{},
	}

	// Print the header and package clause.
	g.Printf("// Code generated by \"%s\"; DO NOT EDIT.\n", command)
	g.Printf("\n")
	g.Printf("package " + "main")
	g.Printf("\n")

	{
		g.Printf("const (\n")

		for directive := range directives {
			switch directive {
			case "page":
				g.Printf("%sIntroStartMarker = `%s`\n", directive, marker(directive, []string{"intro.md"}, true))
				g.Printf("%sIntroEndMarker = `%s`\n", directive, marker(directive, []string{"intro.md"}, false))

				g.Printf("%sFinishStartMarker = `%s`\n", directive, marker(directive, []string{"finish.md"}, true))
				g.Printf("%sFinishEndMarker = `%s`\n", directive, marker(directive, []string{"finish.md"}, false))

			default:
				g.Printf("%sStartMarker = `%s`\n", directive, marker(directive, nil, true))
				g.Printf("%sEndMarker = `%s`\n", directive, marker(directive, nil, false))
			}
		}

		g.Printf(")\n")
	}

	{
		const maxSteps = 20

		g.Printf("var pageStepStartMarkers = [%d]string{\n", maxSteps)

		for i := 1; i < maxSteps; i++ {
			stepFilename := fmt.Sprintf("step%d.md", i)
			g.Printf("\"%s\",\n", marker("page", []string{stepFilename}, true))
		}
		g.Printf("}\n")

		g.Printf("var pageStepEndMarkers = [%d]string{\n", maxSteps)

		for i := 1; i < maxSteps; i++ {
			stepFilename := fmt.Sprintf("step%d.md", i)
			g.Printf("\"%s\",\n", marker("page", []string{stepFilename}, false))
		}
		g.Printf("}\n")
	}

	src, err := g.format()
	if err != nil {
		return fmt.Errorf("couldn't format output: %w", err)
	}

	if err := os.WriteFile(outputPath, src, 0o644); err != nil {
		return fmt.Errorf("couldn't write output file: %w", err)
	}

	return nil
}

type Generator struct {
	buf bytes.Buffer // Accumulated output.
}

func (g *Generator) format() ([]byte, error) {
	return format.Source(g.buf.Bytes())
}

func (g *Generator) Printf(format string, args ...interface{}) {
	fmt.Fprintf(&g.buf, format, args...)
}

const (
	htmlCommentStart = "<!--"
	htmlCommentEnd   = "-->"
	directivePrefix  = "INTERACTIVE"
	directiveStart   = "START"
	directiveEnd     = "END"
)

var directives = map[string][]string{
	"copy":   {},
	"exec":   {},
	"finish": {},
	"ignore": {},
	"intro":  {},
	"page":   {"FILENAME"},
}

func marker(directive string, args []string, isStart bool) string {
	suffix := directiveEnd
	if isStart {
		suffix = directiveStart
	}

	if len(args) == 0 {
		return fmt.Sprintf("%s %s %s %s %s", htmlCommentStart, directivePrefix, directive, suffix, htmlCommentEnd)
	}

	return fmt.Sprintf("%s %s %s %s %s %s", htmlCommentStart, directivePrefix, directive, strings.Join(args, ""), suffix, htmlCommentEnd)
}
