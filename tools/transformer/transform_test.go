package main

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/grafana/killercoda/tools/transformer/goldmark"
	"github.com/grafana/killercoda/tools/transformer/goldmark/renderer/markdown"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

func TestIgnoreTransformer_Transform(t *testing.T) {
	b := &bytes.Buffer{}
	w := bufio.NewWriter(b)
	md := goldmark.NewMarkdown()
	md.Parser().AddOptions(parser.WithASTTransformers(util.Prioritized(&IgnoreTransformer{}, 0)))
	md.SetRenderer(renderer.NewRenderer(renderer.WithNodeRenderers(util.Prioritized(markdown.NewRenderer(), 1000))))

	src := []byte(`## Install Loki and collecting sample logs

<!-- Killercoda ignore START -->

{{< admonition type="note" >}}
This quickstart assumes you are running Linux.
{{< /admonition >}}

<!-- Killercoda ignore END -->

**To install Loki locally, follow these steps:**
`)

	root := md.Parser().Parse(text.NewReader(src))
	require.NoError(t, md.Renderer().Render(w, src, root))

	w.Flush()

	want := `## Install Loki and collecting sample logs

**To install Loki locally, follow these steps:**
`

	assert.Equal(t, want, b.String())
}

func TestIntroTransformer_Transform(t *testing.T) {
	b := &bytes.Buffer{}
	w := bufio.NewWriter(b)
	md := goldmark.NewMarkdown()
	md.Parser().AddOptions(parser.WithASTTransformers(util.Prioritized(&IntroTransformer{}, 0)))
	md.SetRenderer(renderer.NewRenderer(renderer.WithNodeRenderers(util.Prioritized(markdown.NewRenderer(), 1000))))

	src := []byte(`<!-- Killercoda intro.md START -->

# Quickstart to run Loki locally

The Docker Compose configuration instantiates the following components, each in its own container:

- **flog** a sample application which generates log lines.
  [flog](https://github.com/mingrammer/flog) is a log generator for common log formats.
- **Grafana Alloy** which scrapes the log lines from flog, and pushes them to Loki through the gateway.
- **Grafana** which provides visualization of the log lines captured within Loki.

<!-- Killercoda intro.md END -->
`)

	root := md.Parser().Parse(text.NewReader(src))
	require.NoError(t, md.Renderer().Render(w, src, root))

	w.Flush()

	want := `# Quickstart to run Loki locally

The Docker Compose configuration instantiates the following components, each in its own container:

- **flog** a sample application which generates log lines.
  [flog](https://github.com/mingrammer/flog) is a log generator for common log formats.

- **Grafana Alloy** which scrapes the log lines from flog, and pushes them to Loki through the gateway.

- **Grafana** which provides visualization of the log lines captured within Loki.
`

	assert.Equal(t, want, b.String())
}
