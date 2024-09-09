package main

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSubstitutionPreprocessor_Process(t *testing.T) {
	t.Parallel()

	t.Run("remove docs/ignore shortcode", func(t *testing.T) {
		t.Parallel()

		pp := NewSubstitutionPreprocessor(map[*regexp.Regexp][]byte{
			docsIgnoreRegexp: []byte(""),
		})

		src := []byte(`On the Grafana website:

This is rendered before the ignore.

{{< docs/ignore >}}
This isn't rendered.
{{< /docs/ignore >}}

This is rendered after the ignore.
`)

		want := []byte(`On the Grafana website:

This is rendered before the ignore.

This isn't rendered.

This is rendered after the ignore.
`)

		got, err := pp.Process(src)
		require.NoError(t, err)
		assert.Equal(t, string(want), string(got))
	})
}

func TestAdmonitionPreprocessor_Process(t *testing.T) {
	t.Parallel()

	t.Run("Note with one paragraph", func(t *testing.T) {
		t.Parallel()

		pp := NewAdmonitionPreprocessor()
		src := []byte(`{{< admonition type="note" >}}
This quickstart assumes you are running Linux.
{{< /admonition >}}
`)
		want := `> **Note:**
> This quickstart assumes you are running Linux.
`

		got, err := pp.Process(src)
		require.NoError(t, err)
		assert.Equal(t, string(want), string(got))
	})

	t.Run("Note with surrounding nodes", func(t *testing.T) {
		t.Parallel()

		pp := NewAdmonitionPreprocessor()
		src := []byte(`## Install Loki and collecting sample logs

{{< admonition type="note" >}}
This quickstart assumes you are running Linux.
{{< /admonition >}}

**To install Loki locally, follow these steps:**
`)
		want := `## Install Loki and collecting sample logs

> **Note:**
> This quickstart assumes you are running Linux.

**To install Loki locally, follow these steps:**
`

		got, err := pp.Process(src)
		require.NoError(t, err)
		assert.Equal(t, string(want), string(got))
	})

	t.Run("Note with two paragraphs", func(t *testing.T) {
		t.Parallel()

		pp := NewAdmonitionPreprocessor()
		src := []byte(`{{< admonition type="note" >}}
This is a note.

Second paragraph.
{{< /admonition >}}
`)
		want := `> **Note:**
> This is a note.
>
> Second paragraph.
`

		got, err := pp.Process(src)
		require.NoError(t, err)
		assert.Equal(t, string(want), string(got))
	})

	t.Run("Note with extra empty lines", func(t *testing.T) {
		t.Parallel()

		pp := NewAdmonitionPreprocessor()
		src := []byte(`{{< admonition type="note" >}}


This is a note.


Second paragraph.


{{< /admonition >}}
`)
		want := `> **Note:**
> This is a note.
>
> Second paragraph.
`

		got, err := pp.Process(src)
		require.NoError(t, err)
		assert.Equal(t, string(want), string(got))
	})

	t.Run("Indented note", func(t *testing.T) {
		t.Parallel()

		pp := NewAdmonitionPreprocessor()
		src := []byte(`1. First step

   {{< admonition type="note" >}}

   This is a note.

   Second paragraph.

   {{< /admonition >}}

2. Second step
`)
		want := `1. First step

   > **Note:**
   > This is a note.
   >
   > Second paragraph.

2. Second step
`

		got, err := pp.Process(src)
		require.NoError(t, err)
		assert.Equal(t, string(want), string(got))
	})

	t.Run("Note with reference style link", func(t *testing.T) {
		t.Parallel()

		pp := NewAdmonitionPreprocessor()
		src := []byte(`{{< admonition type="tip" >}}
The basic_auth block is commented out because the local docker compose stack doesn't require it.
It's included in this example to show how you can configure authorization for other environments.
For further authorization options, refer to the [loki.write][loki.write] component reference.

[loki.write]: ../../reference/components/loki/loki.write/
{{< /admonition >}}
`)
		want := `> **Tip:**
> The basic_auth block is commented out because the local docker compose stack doesn't require it.
> It's included in this example to show how you can configure authorization for other environments.
> For further authorization options, refer to the [loki.write][loki.write] component reference.

[loki.write]: ../../reference/components/loki/loki.write/
`

		got, err := pp.Process(src)
		require.NoError(t, err)
		assert.Equal(t, string(want), string(got))
	})
}

func TestDocsIgnorePreprocessor_Process(t *testing.T) {
	t.Parallel()

	t.Run("Ignore with one paragraph", func(t *testing.T) {
		t.Parallel()

		pp := NewDocsIgnorePreprocessor()
		src := []byte(`{{< docs/ignore >}}
This isn't rendered in the website but is in Killercoda.
{{< /docs/ignore >}}
`)
		want := `This isn't rendered in the website but is in Killercoda.
`

		got, err := pp.Process(src)
		require.NoError(t, err)
		assert.Equal(t, string(want), string(got))
	})
}
