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
	// referenceLinkDefinitionRegexp is a regular expression that matches single line reference link definitions.
	// https://spec.commonmark.org/0.31.2/#link-reference-definitions
	referenceLinkDefinitionRegexp = regexp.MustCompile(`^\[(?P<label>[^\]]+)\]: (?P<url>.+?)(?: "(?P<title>.+?)")?`)
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
		var (
			matches = admonitionRegexp.FindSubmatch(match)
			body    = matches[admonitionRegexp.SubexpIndex("body")]
			typ     = matches[admonitionRegexp.SubexpIndex("type")]
			// replacement is the buffer that will contain the final replacement.
			// tmp is a buffer used to build the replacement line by line to compress multiple line breaks separating paragraphs into a single one.
			replacement, tmp bytes.Buffer
			// observedIndentation is the number of spaces of indentation observed in most recent line that contained at least one non-whitespace character.
			observedIndentation int
			// newParagraph is a flag that indicates whether the next line should be considered the start of a new paragraph.
			newParagraph = true
			// referenceLinkDefinitions are a set of list destinations found in the admonition body that are to be written after the block quote.
			referenceLinkDefinitions []string
		)

		// > **TYPE:**
		replacement.Write([]byte("> **" + cases.Title(language.English).String(string(typ)) + ":**\n"))

		lines := bytes.Split(body, []byte("\n"))
		for i, line := range lines {
			trimmed := bytes.TrimSpace(line)

			if len(trimmed) > 0 {
				if referenceLinkDefinitionRegexp.Match(trimmed) {
					referenceLinkDefinitions = append(referenceLinkDefinitions, string(trimmed))

					continue
				}

				indent := 0
				for ; indent < len(line) && line[indent] == ' '; indent++ {
					tmp.Write([]byte(" "))
				}
				observedIndentation = indent

				tmp.Write([]byte(">"))
				tmp.Write([]byte(" "))
				tmp.Write(line[indent:])

				if i < len(lines)-1 {
					tmp.Write([]byte("\n"))
				}

				replacement.Write(tmp.Bytes())
				tmp.Reset()

				newParagraph = false

				continue
			}

			if !newParagraph && i < len(lines)-1 {
				for i := 0; i < observedIndentation; i++ {
					tmp.Write([]byte(" "))
				}
				tmp.Write([]byte(">\n"))

				newParagraph = true
			}
		}

		for i, definition := range referenceLinkDefinitions {
			replacement.Write([]byte("\n"))
			replacement.Write([]byte(definition))

			if i == len(referenceLinkDefinitions)-1 {
				replacement.Write([]byte("\n"))
			}
		}

		return replacement.Bytes()
	}), nil
}

type DocsIgnorePreprocessor struct{}

func NewDocsIgnorePreprocessor() *DocsIgnorePreprocessor {
	return &DocsIgnorePreprocessor{}
}

func (ap *DocsIgnorePreprocessor) Process(src []byte) ([]byte, error) {
	return docsIgnoreRegexp.ReplaceAll(src, []byte{}), nil
}
