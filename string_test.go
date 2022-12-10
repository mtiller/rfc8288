package rfc8288

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringify(t *testing.T) {
	assert := assert.New(t)

	entries := []parserTestEntry{
		parseEntry(
			"with href",
			"<https://www.google.com>",
			Link{
				HREF: parseURL("https://www.google.com", t),
			},
		),
		parseEntry(
			"with href, hreflang",
			`<https://www.google.com>; hreflang="en"`,
			Link{
				HREF:     parseURL("https://www.google.com", t),
				HREFLang: "en",
			},
		),
		parseEntry(
			"with href, rel",
			`<https://www.google.com>; rel="next"`,
			Link{
				HREF: parseURL("https://www.google.com", t),
				Rel:  "next",
			},
		),
		parseEntry(
			"with href, media",
			`<https://www.google.com>; media="media"`,
			Link{
				HREF:  parseURL("https://www.google.com", t),
				Media: "media",
			},
		),
		parseEntry(
			"with href, title",
			`<https://www.google.com>; title="title"`,
			Link{
				HREF:  parseURL("https://www.google.com", t),
				Title: "title",
			},
		),
		parseEntry(
			"with href, title*",
			`<https://www.google.com>; title*="title*"`,
			Link{
				HREF:      parseURL("https://www.google.com", t),
				TitleStar: "title*",
			},
		),
		parseEntry(
			"with href, type",
			`<https://www.google.com>; type="type"`,
			Link{
				HREF: parseURL("https://www.google.com", t),
				Type: "type",
			},
		),
	}

	for _, entry := range entries {
		s := entry.result.String()
		assert.Equal(entry.input, s, fmt.Sprintf("test String() output for case '%s'", entry.name))
	}
}