package main

import (
	"fmt"
	"regexp"
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
