package rfc8288

import (
	"encoding/json"
	"net/url"
	"reflect"
	"strings"
)

// UnmarshalJSON unmarshal JSON
func (l *Link) UnmarshalJSON(data []byte) error {

	in := map[string]interface{}{}
	json.Unmarshal(data, &in)

	for k, v := range in {

		switch strings.ToLower(k) {
		case "href":

			if str, ok := v.(string); ok {

				if uri, err := url.Parse(str); err == nil {
					l.HREF = *uri
				} else {

					return &json.UnmarshalTypeError{
						Value:  "uri",
						Type:   reflect.TypeOf(l.HREF),
						Field:  "href",
						Struct: "Link",
					}

				}

			} else {

				return &json.UnmarshalTypeError{
					Value:  "uri",
					Type:   reflect.TypeOf(""),
					Field:  "href",
					Struct: "Link",
				}

			}

		case "rel":

			if str, ok := v.(string); ok {
				l.Rel = str
			} else {
				return &json.UnmarshalTypeError{
					Value:  "string",
					Type:   reflect.TypeOf(""),
					Field:  "rel",
					Struct: "Link",
				}
			}

		case "hreflang":

			if str, ok := v.(string); ok {
				l.HREFLang = str
			} else {
				return &json.UnmarshalTypeError{
					Value:  "string",
					Type:   reflect.TypeOf(int64(0)),
					Field:  "hreflang",
					Struct: "Link",
				}
			}

		case "media":

			if str, ok := v.(string); ok {
				l.Media = str
			} else {
				return &json.UnmarshalTypeError{
					Value:  "string",
					Type:   reflect.TypeOf(""),
					Field:  "media",
					Struct: "Link",
				}
			}

		case "title":

			if str, ok := v.(string); ok {
				l.Title = str
			} else {
				return &json.UnmarshalTypeError{
					Value:  "string",
					Type:   reflect.TypeOf(""),
					Field:  "title",
					Struct: "Link",
				}
			}

		case "title*":

			if str, ok := v.(string); ok {
				l.TitleStar = str
			} else {
				return &json.UnmarshalTypeError{
					Value:  "string",
					Type:   reflect.TypeOf(""),
					Field:  "title*",
					Struct: "Link",
				}
			}

		case "type":

			if str, ok := v.(string); ok {
				l.Type = str
			} else {
				return &json.UnmarshalTypeError{
					Value:  "string",
					Type:   reflect.TypeOf(""),
					Field:  "type",
					Struct: "Link",
				}
			}

		default:

			if err := l.Extend(k, v); err != nil {

				t := reflect.TypeOf(v)

				return &json.UnmarshalTypeError{
					Value:  t.Name(),
					Type:   t,
					Field:  k,
					Struct: "Link",
				}

			}

		}

	}

	return nil

}
