package markdown

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/text"
)

func TestRenderBlockquote(t *testing.T) {
	t.Parallel()

	b := &bytes.Buffer{}
	w := bufio.NewWriter(b)
	md := goldmark.New(goldmark.WithExtensions(NewRenderer()))

	src := []byte(`> **Note:**
> This is an admonition.
`)

	root := md.Parser().Parse(text.NewReader(src))

	require.NoError(t, md.Renderer().Render(w, src, root))

	w.Flush()

	assert.Equal(t, string(src), b.String())
}

func TestRenderCodeblock(t *testing.T) {
	t.Parallel()

	b := &bytes.Buffer{}
	w := bufio.NewWriter(b)
	md := goldmark.New(goldmark.WithExtensions(NewRenderer()))

	src := []byte(`    echo 'Hello, world!'
    echo 'Goodbye, cruel world!'
`)
	root := md.Parser().Parse(text.NewReader(src))

	require.NoError(t, md.Renderer().Render(w, src, root))

	w.Flush()

	assert.Equal(t, string(src), b.String())
}

func TestRenderDocument(t *testing.T) {
	t.Parallel()

	b := &bytes.Buffer{}
	w := bufio.NewWriter(b)
	md := goldmark.New(goldmark.WithExtensions(NewRenderer()))

	src := []byte("# Heading 1\n" +
		"\n" +
		"First paragraph containing some text.\n" +
		"\n" +
		"## Heading 2\n" +
		"\n" +
		"1. An ordered list\n" +
		"\n" +
		"   ```bash\n" +
		"   echo 'Hello, world!'\n" +
		"   ```\n" +
		"\n" +
		"1. of two items\n")
	root := md.Parser().Parse(text.NewReader(src))

	require.NoError(t, md.Renderer().Render(w, src, root))

	w.Flush()

	assert.Equal(t, string(src), b.String())
}

func TestRenderFencedCodeBlock(t *testing.T) {
	t.Parallel()

	b := &bytes.Buffer{}
	w := bufio.NewWriter(b)
	md := goldmark.New(goldmark.WithExtensions(NewRenderer()))

	src := []byte("```go\ngo run ./\n```\n")
	root := md.Parser().Parse(text.NewReader(src))

	require.NoError(t, md.Renderer().Render(w, src, root))

	w.Flush()

	assert.Equal(t, string(src), b.String())
}

func TestRenderHeading(t *testing.T) {
	t.Parallel()

	b := &bytes.Buffer{}
	w := bufio.NewWriter(b)
	md := goldmark.New(goldmark.WithExtensions(NewRenderer()))

	src := []byte("# Heading 1\n")
	root := md.Parser().Parse(text.NewReader(src))

	require.NoError(t, md.Renderer().Render(w, src, root))

	w.Flush()

	assert.Equal(t, string(src), b.String())
}

func TestRenderList(t *testing.T) {
	t.Parallel()

	b := &bytes.Buffer{}
	w := bufio.NewWriter(b)
	md := goldmark.New(goldmark.WithExtensions(NewRenderer()))

	src := []byte(`- One

  - A

- Two

  - B

- Three

  - C
`)
	root := md.Parser().Parse(text.NewReader(src))

	require.NoError(t, md.Renderer().Render(w, src, root))

	w.Flush()

	assert.Equal(t, string(src), b.String())
}

func TestRenderParagraph(t *testing.T) {
	t.Parallel()

	b := &bytes.Buffer{}
	w := bufio.NewWriter(b)
	md := goldmark.New(goldmark.WithExtensions(NewRenderer()))

	src := []byte(`Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.
Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
`)
	root := md.Parser().Parse(text.NewReader(src))

	require.NoError(t, md.Renderer().Render(w, src, root))

	w.Flush()

	assert.Equal(t, string(src), b.String())
}

func TestRenderThematicBreak(t *testing.T) {
	t.Parallel()

	b := &bytes.Buffer{}
	w := bufio.NewWriter(b)
	md := goldmark.New(goldmark.WithExtensions(NewRenderer()))

	src := []byte(`One thing

---

Something else
`)
	root := md.Parser().Parse(text.NewReader(src))

	require.NoError(t, md.Renderer().Render(w, src, root))

	w.Flush()

	assert.Equal(t, string(src), b.String())
}
