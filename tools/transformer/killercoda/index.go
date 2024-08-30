package killercoda

// Populate the Index struct with the parsed data from the Markdown file.
func FromMeta(meta map[any]any) (Index, error) {
	var index Index

	if title, ok := meta["title"].(string); ok {
		index.Title = title
	}

	if description, ok := meta["description"].(string); ok {
		index.Description = description
	}

	if details, ok := meta["details"].(map[any]any); ok {
		if intro, ok := details["intro"].(map[any]any); ok {
			if text, ok := intro["text"].(string); ok {
				index.Details.Intro.Text = text
			}
			if text, ok := intro["foreground"].(string); ok {
				index.Details.Intro.Foreground = text
			}
		}

		if steps, ok := details["steps"].([]any); ok {
			for _, step := range steps {
				if step, ok := step.(map[any]any); ok {
					if text, ok := step["text"].(string); ok {
						index.Details.Steps = append(index.Details.Steps, Text{Text: text})
					}
				}
			}
		}

		if finished, ok := details["finish"].(map[any]any); ok {
			if text, ok := finished["text"].(string); ok {
				index.Details.Finish.Text = text
			}
		}
	}

	if backend, ok := meta["backend"].(map[any]any); ok {
		if imageID, ok := backend["imageid"].(string); ok {
			index.Backend.Imageid = imageID
		}
	}

	return index, nil
}

type Backend struct {
	Imageid string `json:"imageid"`
}

type Details struct {
	Intro  Intro  `json:"intro,omitempty"`
	Steps  []Text `json:"steps"`
	Finish Finish `json:"finish,omitempty"`
}

type Index struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Details     Details `json:"details"`
	Backend     Backend `json:"backend"`
}

type Text struct {
	Text string `json:"text,omitempty"`
}

type Intro struct {
	Text       string `json:"text,omitempty"`
	Foreground string `json:"foreground,omitempty"`
}

type Finish struct {
	Text       string `json:"text,omitempty"`
	Foreground string `json:"foreground,omitempty"`
}
