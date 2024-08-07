//nolint:funlen // Test functions are often lengthy when they include multiple subtests.
package main

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

//nolint:goconst // Strings are repeated for readability.
func TestActionTransformer_Transform(t *testing.T) {
	t.Parallel()

	t.Run("copy directive overrides exec for bash language", func(t *testing.T) {
		t.Parallel()

		b := &bytes.Buffer{}
		w := bufio.NewWriter(b)
		md := goldmark.New(goldmark.WithExtensions(&KillercodaExtension{
			Transformers: []util.PrioritizedValue{},
			AdditionalExtenders: []goldmark.Extender{
				&ActionTransformer{Kind: "copy"},
				&ActionTransformer{Kind: "exec"},
			},
		}))

		src := []byte("1. Create a directory called `evaluate-loki` for the demo environment.\n" +
			"   Make `evaluate-loki` your current working directory:\n" +
			"\n" +
			"   <!-- INTERACTIVE copy START -->\n" +
			"\n" +
			"   ```bash\n" +
			"   mkdir evaluate-loki\n" +
			"   cd evaluate-loki\n" +
			"   ```\n" +
			"\n" +
			"   <!-- INTERACTIVE copy END -->\n")

		root := md.Parser().Parse(text.NewReader(src))
		require.NoError(t, md.Renderer().Render(w, src, root))

		w.Flush()

		want := "1. Create a directory called `evaluate-loki` for the demo environment.\n" +
			"   Make `evaluate-loki` your current working directory:\n" +
			"\n" +
			"   ```bash\n" +
			"   mkdir evaluate-loki\n" +
			"   cd evaluate-loki\n" +
			"   ```{{copy}}\n"

		assert.Equal(t, want, b.String())
	})

	t.Run("bash language defaults to exec", func(t *testing.T) {
		t.Parallel()

		b := &bytes.Buffer{}
		w := bufio.NewWriter(b)
		md := goldmark.New(goldmark.WithExtensions(&KillercodaExtension{
			Transformers: []util.PrioritizedValue{},
			AdditionalExtenders: []goldmark.Extender{
				&ActionTransformer{Kind: "copy"},
				&ActionTransformer{Kind: "exec"},
			},
		}))

		src := []byte("1. Create a directory called `evaluate-loki` for the demo environment.\n" +
			"   Make `evaluate-loki` your current working directory:\n" +
			"\n" +
			"   ```bash\n" +
			"   mkdir evaluate-loki\n" +
			"   cd evaluate-loki\n" +
			"   ```\n")

		root := md.Parser().Parse(text.NewReader(src))
		require.NoError(t, md.Renderer().Render(w, src, root))

		w.Flush()

		want := "1. Create a directory called `evaluate-loki` for the demo environment.\n" +
			"   Make `evaluate-loki` your current working directory:\n" +
			"\n" +
			"   ```bash\n" +
			"   mkdir evaluate-loki\n" +
			"   cd evaluate-loki\n" +
			"   ```{{exec}}\n"

		assert.Equal(t, want, b.String())
	})

	t.Run("exec directive overrides copy default for other languages", func(t *testing.T) {
		t.Parallel()

		b := &bytes.Buffer{}
		w := bufio.NewWriter(b)
		md := goldmark.New(goldmark.WithExtensions(&KillercodaExtension{
			Transformers: []util.PrioritizedValue{},
			AdditionalExtenders: []goldmark.Extender{
				&ActionTransformer{Kind: "copy"},
				&ActionTransformer{Kind: "exec"},
			},
		}))

		src := []byte("1. Create a directory called `evaluate-loki` for the demo environment.\n" +
			"   Make `evaluate-loki` your current working directory:\n" +
			"\n" +
			"   <!-- INTERACTIVE exec START -->\n" +
			"\n" +
			"   ```\n" +
			"   mkdir evaluate-loki\n" +
			"   cd evaluate-loki\n" +
			"   ```\n" +
			"\n" +
			"   <!-- INTERACTIVE exec END -->\n")

		root := md.Parser().Parse(text.NewReader(src))
		require.NoError(t, md.Renderer().Render(w, src, root))

		w.Flush()

		want := "1. Create a directory called `evaluate-loki` for the demo environment.\n" +
			"   Make `evaluate-loki` your current working directory:\n" +
			"\n" +
			"   ```\n" +
			"   mkdir evaluate-loki\n" +
			"   cd evaluate-loki\n" +
			"   ```{{exec}}\n"

		assert.Equal(t, want, b.String())
	})

	t.Run("other languages default to copy", func(t *testing.T) {
		t.Parallel()

		b := &bytes.Buffer{}
		w := bufio.NewWriter(b)
		md := goldmark.New(goldmark.WithExtensions(&KillercodaExtension{
			Transformers: []util.PrioritizedValue{},
			AdditionalExtenders: []goldmark.Extender{
				&ActionTransformer{Kind: "copy"},
				&ActionTransformer{Kind: "exec"},
			},
		}))

		src := []byte("1. Create a directory called `evaluate-loki` for the demo environment.\n" +
			"   Make `evaluate-loki` your current working directory:\n" +
			"\n" +
			"   ```\n" +
			"   mkdir evaluate-loki\n" +
			"   cd evaluate-loki\n" +
			"   ```\n")

		root := md.Parser().Parse(text.NewReader(src))
		require.NoError(t, md.Renderer().Render(w, src, root))

		w.Flush()

		want := "1. Create a directory called `evaluate-loki` for the demo environment.\n" +
			"   Make `evaluate-loki` your current working directory:\n" +
			"\n" +
			"   ```\n" +
			"   mkdir evaluate-loki\n" +
			"   cd evaluate-loki\n" +
			"   ```{{copy}}\n"

		assert.Equal(t, want, b.String())
	})
}

func TestAdmonitionTransformer_Transform(t *testing.T) {
	t.Skip("not implemented")
	t.Parallel()

	t.Run("with {{< syntax in a single paragraph", func(t *testing.T) {
		t.Parallel()

		b := &bytes.Buffer{}
		w := bufio.NewWriter(b)
		md := goldmark.New(goldmark.WithExtensions(&KillercodaExtension{
			Transformers: []util.PrioritizedValue{},
			AdditionalExtenders: []goldmark.Extender{
				&AdmonitionTransformer{},
			},
		}))

		src := []byte(`{{< admonition type="note" >}}
This is a note.
{{< /admonition >}}
`)

		root := md.Parser().Parse(text.NewReader(src))
		require.NoError(t, md.Renderer().Render(w, src, root))

		w.Flush()

		want := `> **Note:**
> This is a note.
`

		assert.Equal(t, want, b.String())
	})
	t.Run("with {{< syntax over multiple paragraphs", func(t *testing.T) {
		t.Parallel()

		b := &bytes.Buffer{}
		w := bufio.NewWriter(b)
		md := goldmark.New(goldmark.WithExtensions(&KillercodaExtension{
			Transformers: []util.PrioritizedValue{},
			AdditionalExtenders: []goldmark.Extender{
				&AdmonitionTransformer{},
			},
		}))

		src := []byte(`{{< admonition type="note" >}}

This is a note.

{{< /admonition >}}
`)

		root := md.Parser().Parse(text.NewReader(src))
		require.NoError(t, md.Renderer().Render(w, src, root))

		w.Flush()

		want := `> **Note:**
> This is a note.
`

		assert.Equal(t, want, b.String())
	})
}

func TestFigureTransformer_Transform(t *testing.T) {
	t.Parallel()

	t.Run("with alt argument", func(t *testing.T) {
		t.Parallel()

		b := &bytes.Buffer{}
		w := bufio.NewWriter(b)
		md := goldmark.New(goldmark.WithExtensions(&KillercodaExtension{
			Transformers: []util.PrioritizedValue{},
			AdditionalExtenders: []goldmark.Extender{
				&FigureTransformer{},
			},
		}))

		src := []byte("{{< figure src=\"/media/docs/loki/grafana-query-builder-v2.png\" caption=\"Grafana Explore\" alt=\"Grafana Explore\" >}}\n")

		root := md.Parser().Parse(text.NewReader(src))
		require.NoError(t, md.Renderer().Render(w, src, root))

		w.Flush()

		want := "![Grafana Explore](/media/docs/loki/grafana-query-builder-v2.png)\n"

		assert.Equal(t, want, b.String())
	})

	t.Run("with caption instead of alt", func(t *testing.T) {
		t.Parallel()

		b := &bytes.Buffer{}
		w := bufio.NewWriter(b)
		md := goldmark.New(goldmark.WithExtensions(&KillercodaExtension{
			Transformers: []util.PrioritizedValue{},
			AdditionalExtenders: []goldmark.Extender{
				&FigureTransformer{},
			},
		}))

		src := []byte("{{< figure src=\"/media/docs/loki/grafana-query-builder-v2.png\" caption=\"Grafana Explore\" >}}\n")

		root := md.Parser().Parse(text.NewReader(src))
		require.NoError(t, md.Renderer().Render(w, src, root))

		w.Flush()

		want := "![Grafana Explore](/media/docs/loki/grafana-query-builder-v2.png)\n"

		assert.Equal(t, want, b.String())
	})

	t.Run("no alt", func(t *testing.T) {
		t.Parallel()

		b := &bytes.Buffer{}
		w := bufio.NewWriter(b)
		md := goldmark.New(goldmark.WithExtensions(&KillercodaExtension{
			Transformers: []util.PrioritizedValue{},
			AdditionalExtenders: []goldmark.Extender{
				&FigureTransformer{},
			},
		}))

		src := []byte("{{< figure src=\"/media/docs/loki/grafana-query-builder-v2.png\" >}}\n")

		root := md.Parser().Parse(text.NewReader(src))
		require.NoError(t, md.Renderer().Render(w, src, root))

		w.Flush()

		want := "![](/media/docs/loki/grafana-query-builder-v2.png)\n"

		assert.Equal(t, want, b.String())
	})
}

func TestIgnoreTransformer_Transform(t *testing.T) {
	t.Parallel()

	b := &bytes.Buffer{}
	w := bufio.NewWriter(b)
	md := goldmark.New(goldmark.WithExtensions(&KillercodaExtension{
		Transformers: []util.PrioritizedValue{},
		AdditionalExtenders: []goldmark.Extender{
			&IgnoreTransformer{},
		},
	}))

	src := []byte(`## Install Loki and collecting sample logs

<!-- INTERACTIVE ignore START -->

{{< admonition type="note" >}}
This quickstart assumes you are running Linux.
{{< /admonition >}}

<!-- INTERACTIVE ignore END -->

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

func TestLinkTransformer_Transform(t *testing.T) {
	t.Parallel()

	b := &bytes.Buffer{}
	w := bufio.NewWriter(b)
	md := goldmark.New(goldmark.WithExtensions(&KillercodaExtension{
		Transformers: []util.PrioritizedValue{},
		AdditionalExtenders: []goldmark.Extender{
			&LinkTransformer{},
		},
	}))

	src := []byte(`You can view your logs using the command line interface, [LogCLI](/docs/loki/latest/query/logcli/), but the easiest way to view your logs is with Grafana.

If you want to experiment with Loki, you can run Loki locally using the Docker Compose file that ships with Loki.
It runs Loki in a [monolithic deployment](https://grafana.com/docs/loki/<LOKI_VERSION>/get-started/deployment-modes/#monolithic-mode) mode and includes a sample application to generate logs.

- You can access the Grafana Alloy UI at [http://localhost:12345/ready](http://localhost:12345/ready).
`)

	root := md.Parser().Parse(text.NewReader(src))
	require.NoError(t, md.Renderer().Render(w, src, root))

	w.Flush()

	want := `You can view your logs using the command line interface, [LogCLI](https://grafana.com/docs/loki/latest/query/logcli/), but the easiest way to view your logs is with Grafana.

If you want to experiment with Loki, you can run Loki locally using the Docker Compose file that ships with Loki.
It runs Loki in a [monolithic deployment](https://grafana.com/docs/loki/latest/get-started/deployment-modes/#monolithic-mode) mode and includes a sample application to generate logs.

- You can access the Grafana Alloy UI at [http://localhost:12345/ready]({{TRAFFIC_HOST1_12345}}/ready).
`

	assert.Equal(t, want, b.String())
}

func TestStepTransformer_Transform(t *testing.T) {
	t.Parallel()

	t.Run("intro", func(t *testing.T) {
		t.Parallel()

		b := &bytes.Buffer{}
		w := bufio.NewWriter(b)

		md := goldmark.New(goldmark.WithExtensions(&KillercodaExtension{
			Transformers: []util.PrioritizedValue{},
			AdditionalExtenders: []goldmark.Extender{
				&StepTransformer{StartMarker: pageIntroStartMarker, EndMarker: pageIntroEndMarker},
			},
		}))

		src := []byte(`<!-- INTERACTIVE page intro.md START -->

# Quickstart to run Loki locally

The Docker Compose configuration instantiates the following components, each in its own container:

- **flog** a sample application which generates log lines.
  [flog](https://github.com/mingrammer/flog) is a log generator for common log formats.
- **Grafana Alloy** which scrapes the log lines from flog, and pushes them to Loki through the gateway.
- **Grafana** which provides visualization of the log lines captured within Loki.

<!-- INTERACTIVE page intro.md END -->
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
	})
}
