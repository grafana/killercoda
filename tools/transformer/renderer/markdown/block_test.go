package markdown

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/grafana/killercoda/tools/transformer/parser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

func TestRenderDocument(t *testing.T) {
	t.Parallel()

	b := &bytes.Buffer{}
	w := bufio.NewWriter(b)
	r := renderer.NewRenderer(renderer.WithNodeRenderers(util.Prioritized(NewRenderer(), 1000)))

	src := []byte(`# Heading 1

First paragraph containing some text.

## Heading 2

- An unordered list
- of two items
`)
	root := parser.New().Parse(text.NewReader(src))

	require.NoError(t, r.Render(w, src, root))

	w.Flush()

	assert.Equal(t, string(src), b.String())
}

func TestRenderFencedCodeBlock(t *testing.T) {
	t.Parallel()

	b := &bytes.Buffer{}
	w := bufio.NewWriter(b)
	r := renderer.NewRenderer(renderer.WithNodeRenderers(util.Prioritized(NewRenderer(), 1000)))

	src := []byte("```go\ngo run ./\n```\n")
	root := parser.New().Parse(text.NewReader(src))

	require.NoError(t, r.Render(w, src, root))

	w.Flush()

	assert.Equal(t, string(src), b.String())
}

func TestRenderHeading(t *testing.T) {
	t.Parallel()

	b := &bytes.Buffer{}
	w := bufio.NewWriter(b)
	r := renderer.NewRenderer(renderer.WithNodeRenderers(util.Prioritized(NewRenderer(), 1000)))

	src := []byte("# Heading 1\n")
	root := parser.New().Parse(text.NewReader(src))

	require.NoError(t, r.Render(w, src, root))

	w.Flush()

	assert.Equal(t, string(src), b.String())
}

func TestRenderList(t *testing.T) {
	t.Parallel()

	b := &bytes.Buffer{}
	w := bufio.NewWriter(b)
	r := renderer.NewRenderer(renderer.WithNodeRenderers(util.Prioritized(NewRenderer(), 1000)))

	src := []byte(`- One
  - A
- Two
  - B
- Three
  - C
- Loose
  - D
`)
	root := parser.New().Parse(text.NewReader(src))

	require.NoError(t, r.Render(w, src, root))

	w.Flush()

	assert.Equal(t, string(src), b.String())
}

func TestRenderParagraph(t *testing.T) {
	t.Parallel()

	b := &bytes.Buffer{}
	w := bufio.NewWriter(b)
	r := renderer.NewRenderer(renderer.WithNodeRenderers(util.Prioritized(NewRenderer(), 1000)))

	src := []byte(`Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.
Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
`)
	root := parser.New().Parse(text.NewReader(src))

	require.NoError(t, r.Render(w, src, root))

	w.Flush()

	assert.Equal(t, string(src), b.String())
}
