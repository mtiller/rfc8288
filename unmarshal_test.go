package rfc8288

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type unmarshalTestEntry struct {
	name string
	m    map[string]interface{}
	link Link
}

func unmarshalEntry(name string, m map[string]interface{}, link Link) unmarshalTestEntry {
	return unmarshalTestEntry{
		name: name,
		m:    m,
		link: link,
	}
}

func TestUnmarshal(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)

	entries := []unmarshalTestEntry{
		unmarshalEntry(
			"should unmarshal all fields and extensions",
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
		),
	}

	for _, entry := range entries {
		in := entry.m
		out := entry.link
		// given
		jsonBytes, err := json.Marshal(in)
		require.NoError(err)
		result := Link{}

		// when
		json.Unmarshal(jsonBytes, &result)

		// then
		if _, ok := in["href"]; ok {
			assert.Equal(out.HREF.String(), result.HREF.String())
		}

		if _, ok := in["rel"]; ok {
			assert.Equal(out.Rel, result.Rel)
		}

		if _, ok := in["hreflang"]; ok {
			assert.Equal(out.HREFLang, result.HREFLang)
		}

		if _, ok := in["media"]; ok {
			assert.Equal(out.Media, result.Media)
		}

		if _, ok := in["title"]; ok {
			assert.Equal(out.Title, result.Title)
		}

		if _, ok := in["title*"]; ok {
			assert.Equal(out.TitleStar, result.TitleStar)
		}

		if _, ok := in["type"]; ok {
			assert.Equal(out.Type, result.Type)
		}

		for key := range in {
			if _, isReserved := ReservedKeys[key]; isReserved {
				continue
			}

			resultValue, resultExists := result.Extension(key)
			assert.True(resultExists)

			outValue, valueExists := out.Extension(key)
			assert.True(valueExists)

			assert.Equal(outValue, resultValue)
		}
	}
}
