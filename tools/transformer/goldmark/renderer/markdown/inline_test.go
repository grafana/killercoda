package markdown

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/grafana/killercoda/tools/transformer/goldmark"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

func TestRenderCodespan(t *testing.T) {
	t.Parallel()

	b := &bytes.Buffer{}
	w := bufio.NewWriter(b)
	md := goldmark.NewMarkdown()
	md.SetRenderer(renderer.NewRenderer(renderer.WithNodeRenderers(util.Prioritized(NewRenderer(), 1000))))

	src := []byte("`code`\n")
	root := md.Parser().Parse(text.NewReader(src))

	require.NoError(t, md.Renderer().Render(w, src, root))

	w.Flush()

	assert.Equal(t, string(src), b.String())
}

func TestRenderEmphasis(t *testing.T) {
	t.Parallel()

	b := &bytes.Buffer{}
	w := bufio.NewWriter(b)
	md := goldmark.NewMarkdown()
	md.SetRenderer(renderer.NewRenderer(renderer.WithNodeRenderers(util.Prioritized(NewRenderer(), 1000))))

	src := []byte("**Strong** and _italic_\n")
	root := md.Parser().Parse(text.NewReader(src))

	require.NoError(t, md.Renderer().Render(w, src, root))

	w.Flush()

	assert.Equal(t, string(src), b.String())
}

func TestRenderLink(t *testing.T) {
	t.Parallel()

	b := &bytes.Buffer{}
	w := bufio.NewWriter(b)
	md := goldmark.NewMarkdown()
	md.SetRenderer(renderer.NewRenderer(renderer.WithNodeRenderers(util.Prioritized(NewRenderer(), 1000))))

	src := []byte("[TEXT](DESTINATION)\n")
	root := md.Parser().Parse(text.NewReader(src))

	require.NoError(t, md.Renderer().Render(w, src, root))

	w.Flush()

	assert.Equal(t, string(src), b.String())
}

func TestRenderText(t *testing.T) {
	t.Parallel()

	b := &bytes.Buffer{}
	w := bufio.NewWriter(b)
	md := goldmark.NewMarkdown()
	md.SetRenderer(renderer.NewRenderer(renderer.WithNodeRenderers(util.Prioritized(NewRenderer(), 1000))))

	src := []byte("'<GRAFANA_VERSION>'\n")
	root := md.Parser().Parse(text.NewReader(src))

	root.Dump(src, 0)

	require.NoError(t, md.Renderer().Render(w, src, root))

	w.Flush()

	// TODO: avoid smart quotes?
	assert.Equal(t, "‘<GRAFANA_VERSION>’\n", b.String())
}
