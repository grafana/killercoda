package markdown

func TestRenderCodespan(t *testing.T) {
	t.Parallel()

	b := &bytes.Buffer{}
	w := bufio.NewWriter(b)
	r := renderer.NewRenderer(renderer.WithNodeRenderers(util.Prioritized(NewRenderer(), 1000)))

	src := []byte("`code`\n")
	root := parser.New().Parse(text.NewReader(src))

	require.NoError(t, r.Render(w, src, root))

	w.Flush()

	assert.Equal(t, string(src), b.String())
}

func TestRenderEmphasis(t *testing.T) {
	t.Parallel()

	b := &bytes.Buffer{}
	w := bufio.NewWriter(b)
	r := renderer.NewRenderer(renderer.WithNodeRenderers(util.Prioritized(NewRenderer(), 1000)))

	src := []byte("**Strong** and _italic_\n")
	root := parser.New().Parse(text.NewReader(src))

	require.NoError(t, r.Render(w, src, root))

	w.Flush()

	assert.Equal(t, string(src), b.String())
}

func TestRenderLink(t *testing.T) {
	t.Parallel()

	b := &bytes.Buffer{}
	w := bufio.NewWriter(b)
	r := renderer.NewRenderer(renderer.WithNodeRenderers(util.Prioritized(NewRenderer(), 1000)))

	src := []byte("[TEXT](DESTINATION)\n")
	root := parser.New().Parse(text.NewReader(src))

	require.NoError(t, r.Render(w, src, root))

	w.Flush()

	assert.Equal(t, string(src), b.String())
}
