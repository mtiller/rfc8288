package rfc8288

// ParseLink attempts to parse a link string
func ParseLink(link string) (Link, error) {

	var (
		rs io.RuneScanner = strings.NewReader(link)
		s                 = scanner{runeScanner: rs}
		p                 = parser{scanner: s}
	)

	return p.parse()

}

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type parserTestEntry struct {
	name   string
	input  string
	result Link
}

func parseEntry(name, input string, result Link) parserTestEntry {
	return parserTestEntry{
		name:   name,
		input:  input,
		result: result,
	}
}

func TestParser(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	entries := []parserTestEntry{
		parseEntry(
			"href",
			`<https://www.google.com>`,
			Link{
				HREF: parseURL("https://www.google.com", t),
			},
		),
		parseEntry(
			"href, rel",
			`<https://www.google.com>; rel="next"`,
			Link{
				HREF: parseURL("https://www.google.com", t),
				Rel:  "next",
			},
		),
		parseEntry(
			"href, rel, hreflang",
			`<https://www.google.com>; rel="next"; hreflang="en"`,
			Link{
				HREF:     parseURL("https://www.google.com", t),
				Rel:      "next",
				HREFLang: "en",
			},
		),
		parseEntry(
			"href, rel, hreflang, media",
			`<https://www.google.com>; rel="next"; hreflang="en"; media="media"`,
			Link{
				HREF:     parseURL("https://www.google.com", t),
				Rel:      "next",
				HREFLang: "en",
				Media:    "media",
			},
		),
		parseEntry(
			"href, rel, hreflang, title",
			`<https://www.google.com>; rel="next"; hreflang="en"; title="title"`,
			Link{
				HREF:     parseURL("https://www.google.com", t),
				Rel:      "next",
				HREFLang: "en",
				Title:    "title",
			},
		),
		parseEntry(
			"href, rel, hreflang, title, title*",
			`<https://www.google.com>; rel="next"; hreflang="en"; title="title"; title*="title*"`,
			Link{
				HREF:      parseURL("https://www.google.com", t),
				Rel:       "next",
				HREFLang:  "en",
				Title:     "title",
				TitleStar: "title*",
			},
		),
		parseEntry(
			"href, rel, hreflang, title, title*, type",
			`<https://www.google.com>; rel="next"; hreflang="en"; title="title"; title*="title*"; type="type"`,
			Link{
				HREF:      parseURL("https://www.google.com", t),
				Rel:       "next",
				HREFLang:  "en",
				Title:     "title",
				TitleStar: "title*",
				Type:      "type",
			},
		),
		parseEntry(
			"href, rel, hreflang, title, title*, type, extensions",
			`<https://www.google.com>; rel="next"; hreflang="en"; title="title"; title*="title*"; type="type"; extension="value"`,
			Link{
				HREF:          parseURL("https://www.google.com", t),
				Rel:           "next",
				HREFLang:      "en",
				Title:         "title",
				TitleStar:     "title*",
				Type:          "type",
				extensionKeys: []string{"extension"},
				extensions: map[string]interface{}{
					"extension": "value",
				},
			},
		),
	}

	for i, entry := range entries {
		r, err := ParseLink(entry.input)
		require.NoError(err, fmt.Sprintf("Error parsing entry %d (%s)", i, entry.input))
		assert.Equal(entry.result, r, fmt.Sprintf("Mismatch for parser entry %d (%s)", i, entry.name))
	}
}
