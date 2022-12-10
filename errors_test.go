package rfc8288

import (
	"encoding/json"
	"net/url"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type errorUnmarshalEntry struct {
	name  string
	input string
	err   error
}

func errorEntry(name, input string, err error) errorUnmarshalEntry {
	return errorUnmarshalEntry{
		name:  name,
		input: input,
		err:   err,
	}
}

func TestUnmarshalErrors(t *testing.T) {
	assert := assert.New(t)
	entries := []errorUnmarshalEntry{
		errorEntry(
			"should return json.UnmarshalTypeError describing href field",
			`{
				"href": "Not a valid url !@#$%^&*()_+"
			}`,
			&json.UnmarshalTypeError{
				Value:  "uri",
				Type:   reflect.TypeOf(url.URL{}),
				Field:  "href",
				Struct: "Link",
			},
		),
		errorEntry(
			"should return json.UnmarshalTypeError describing href field",
			`{
				"href": false
			}`,
			&json.UnmarshalTypeError{
				Value:  "uri",
				Type:   reflect.TypeOf(""),
				Field:  "href",
				Struct: "Link",
			},
		),
		errorEntry(
			"should return json.UnmarshalTypeError describing rel field",
			`{
				"rel": false
			}`,
			&json.UnmarshalTypeError{
				Value:  "string",
				Type:   reflect.TypeOf(""),
				Field:  "rel",
				Struct: "Link",
			},
		),
		errorEntry(
			"should return json.UnmarshalTypeError describing hreflang field",
			`{
				"hreflang": false
			}`,
			&json.UnmarshalTypeError{
				Value:  "string",
				Type:   reflect.TypeOf(int64(0)),
				Field:  "hreflang",
				Struct: "Link",
			},
		),
		errorEntry(
			"should return json.UnmarshalTypeError describing media field",
			`{
				"media": false
			}`,
			&json.UnmarshalTypeError{
				Value:  "string",
				Type:   reflect.TypeOf(""),
				Field:  "media",
				Struct: "Link",
			},
		),
		errorEntry(
			"should return json.UnmarshalTypeError describing title field",
			`{
				"title": false
			}`,
			&json.UnmarshalTypeError{
				Value:  "string",
				Type:   reflect.TypeOf(""),
				Field:  "title",
				Struct: "Link",
			},
		),
		errorEntry(
			"should return json.UnmarshalTypeError describing title* field",
			`{
				"title*": false
			}`,
			&json.UnmarshalTypeError{
				Value:  "string",
				Type:   reflect.TypeOf(""),
				Field:  "title*",
				Struct: "Link",
			},
		),
		errorEntry(
			"should return json.UnmarshalTypeError describing type field",
			`{
				"type": false
			}`,
			&json.UnmarshalTypeError{
				Value:  "string",
				Type:   reflect.TypeOf(""),
				Field:  "type",
				Struct: "Link",
			},
		),
	}

	for _, entry := range entries {
		l := Link{}
		err := json.Unmarshal([]byte(entry.input), &l)
		assert.Equal(entry.err, err)
	}
}
