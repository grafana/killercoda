package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWriteFileRejectsMissingEndMarker(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "finish.md")
	require.NoError(t, os.WriteFile(path, []byte("existing tutorial"), 0600))
	md := NewTransformerMarkdown(DefaultKillercodaTransformers,
		&StepTransformer{StartMarker: pageFinishStartMarker, EndMarker: pageFinishEndMarker})
	source := []byte(pageFinishStartMarker + "\n\nNext steps\n")
	require.ErrorIs(t, writeFile(md, dir, "finish.md", source), errNoEndMarker)
	content, err := os.ReadFile(path)
	require.NoError(t, err)
	require.Equal(t, "existing tutorial", string(content))

	// A successful parse clears errors from a previous document.
	source = append(source, []byte("\n"+pageFinishEndMarker+"\n")...)
	require.NoError(t, writeFile(md, dir, "finish.md", source))
	content, err = os.ReadFile(path)
	require.NoError(t, err)
	require.Contains(t, string(content), "Next steps")
	require.NotContains(t, string(content), "INTERACTIVE")
}
