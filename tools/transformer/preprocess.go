package main

import (
	"bytes"
	"fmt"
	"regexp"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	admonitionRegexp = regexp.MustCompile(`(?s){{[<%] * admonition +type="(?P<type>[^"]+)" *[%>]}}(?P<body>.+?){{[<%] */admonition *[%>]}}\n?`)
	docsIgnoreRegexp = regexp.MustCompile(`{{< *?/?docs/ignore *?>}}\n?`)
)

// Preprocessor processes raw Markdown source file bytes.
type Preprocessor interface {
	// Process processes raw Markdown source file bytes.
	Process(src []byte) ([]byte, error)
}

type ComposedPreprocessor struct {
	preprocessors []Preprocessor
}

func NewComposedPreprocessor(preprocessors ...Preprocessor) *ComposedPreprocessor {
	return &ComposedPreprocessor{preprocessors: preprocessors}
}

func (cp *ComposedPreprocessor) Process(src []byte) ([]byte, error) {
	var err error
	for _, p := range cp.preprocessors {
		src, err = p.Process(src)
		if err != nil {
			return nil, err
		}
	}

	return src, nil
}

type SubstitutionPreprocessor struct {
	substitutions map[*regexp.Regexp][]byte
}

func NewSubstitutionPreprocessor(substitutions map[*regexp.Regexp][]byte) *SubstitutionPreprocessor {
	return &SubstitutionPreprocessor{substitutions: substitutions}
}

func NewSubstitutionPreprocessorFromMeta(meta map[any]any) (*SubstitutionPreprocessor, error) {
	subs := make(map[*regexp.Regexp][]byte)

	if preprocessing, ok := meta["preprocessing"].(map[any]any); ok {
		if substitutions, ok := preprocessing["substitutions"].([]any); ok {
			for _, substitution := range substitutions {
				if s, ok := substitution.(map[any]any); ok {
					if expr, ok := s["regexp"].(string); ok {
						if replacement, ok := s["replacement"].(string); ok {
							re, err := regexp.Compile(expr)
							if err != nil {
								return nil, fmt.Errorf("couldn't compile regular expression %q: %w", expr, err)
							}

							subs[re] = []byte(replacement)
						}
					}
				}
			}
		}
	}

	return &SubstitutionPreprocessor{substitutions: subs}, nil
}

func (sp *SubstitutionPreprocessor) AddSubstitution(re *regexp.Regexp, replacement []byte) error {
	sp.substitutions[re] = replacement

	return nil
}

func (sp *SubstitutionPreprocessor) Process(src []byte) ([]byte, error) {
	for expr, replacement := range sp.substitutions {
		src = expr.ReplaceAll(src, []byte(replacement))
	}

	return src, nil
}

type AdmonitionPreprocessor struct{}

func NewAdmonitionPreprocessor() *AdmonitionPreprocessor {
	return &AdmonitionPreprocessor{}
}

func (ap *AdmonitionPreprocessor) Process(src []byte) ([]byte, error) {
	return admonitionRegexp.ReplaceAllFunc(src, func(match []byte) []byte {
		matches := admonitionRegexp.FindSubmatch(match)
		typ := matches[admonitionRegexp.SubexpIndex("type")]
		body := matches[admonitionRegexp.SubexpIndex("body")]

		var b, tmp bytes.Buffer
		b.Write([]byte("> **" + cases.Title(language.English).String(string(typ)) + ":**\n"))

		defaultIndent := 0
		newParagraph := true

		lines := bytes.Split(body, []byte("\n"))
		for i, line := range lines {
			trimmed := bytes.TrimSpace(line)

			if len(trimmed) > 0 {
				indent := 0
				for ; indent < len(line) && line[indent] == ' '; indent++ {
					tmp.Write([]byte(" "))
				}
				defaultIndent = indent

				tmp.Write([]byte(">"))
				tmp.Write([]byte(" "))
				tmp.Write(line[indent:])

				if i < len(lines)-1 {
					tmp.Write([]byte("\n"))
				}

				b.Write(tmp.Bytes())
				tmp.Reset()

				newParagraph = false

				continue
			}

			if !newParagraph && i < len(lines)-1 {
				for i := 0; i < defaultIndent; i++ {
					tmp.Write([]byte(" "))
				}
				tmp.Write([]byte(">\n"))

				newParagraph = true
			}
		}

		return b.Bytes()
	}), nil
}

type DocsIgnorePreprocessor struct{}

func NewDocsIgnorePreprocessor() *DocsIgnorePreprocessor {
	return &DocsIgnorePreprocessor{}
}

func (ap *DocsIgnorePreprocessor) Process(src []byte) ([]byte, error) {
	return docsIgnoreRegexp.ReplaceAll(src, []byte{}), nil
}
