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
