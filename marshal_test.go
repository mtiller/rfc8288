package rfc8288

import (
	"encoding/json"
	"fmt"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func parseURL(u string, t *testing.T) url.URL {
	require := require.New(t)

	uri, err := url.Parse(u)
	require.NoError(err)

	return *uri
}

type marshalTestEntry struct {
	name string
	link Link
	m    map[string]interface{}
}

func marshalEntry(name string, link Link, m map[string]interface{}) marshalTestEntry {
	return marshalTestEntry{
		name: name,
		link: link,
		m:    m,
	}
}

func TestLinkMarshal(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	entries := []marshalTestEntry{
		marshalEntry(
			"should marshal all fields and extensions",
			func() Link {

				l := Link{
					HREF:      parseURL("https://www.google.com", t),
					Rel:       "rel",
					HREFLang:  "hreflang",
					Media:     "media",
					Title:     "title",
					TitleStar: "title*",
					Type:      "type",
				}

				l.Extend("extension", "value")

				return l

			}(),
			map[string]interface{}{
				"href":      "https://www.google.com",
				"rel":       "rel",
				"hreflang":  "hreflang",
				"media":     "media",
				"title":     "title",
				"title*":    "title*",
				"type":      "type",
				"extension": "value",
			},
		),
	}
	for _, entry := range entries {
		in := entry.link
		out := entry.m
		// given
		jsonBytes, err := json.Marshal(entry.link)
		result := make(map[string]interface{})

		require.NoError(err)

		// when
		json.Unmarshal(jsonBytes, &result)

		check := func(prop string, entry marshalTestEntry) {
			assert.Equal(out[prop], result[prop], fmt.Sprintf("Comparing %s property for case %s", prop, entry.name))
		}

		// then
		var zero url.URL
		if in.HREF != zero {
			check("href", entry)
		}

		if in.Rel != "" {
			check("rel", entry)
		}

		if in.HREFLang != "" {
			check("hreflang", entry)
		}

		if in.Media != "" {
			check("media", entry)
		}

		if in.Title != "" {
			check("title", entry)
		}

		if in.TitleStar != "" {
			check("title*", entry)
		}

		if in.Type != "" {
			check("type", entry)
		}

		for _, key := range in.ExtensionKeys() {
			check(key, entry)
		}
	}
}
