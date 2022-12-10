package rfc8288

import (
	"encoding/json"
	"net/url"
)

// MarshalJSON Marshals JSON
func (l Link) MarshalJSON() ([]byte, error) {

	out := map[string]interface{}{}

	var zero url.URL
	if l.HREF != zero {
		out["href"] = l.HREF.String()
	}

	if l.Rel != "" {
		out["rel"] = l.Rel
	}

	if l.Rev != "" {
		out["rev"] = l.Rev
	}

	if l.Anchor != "" {
		out["anchor"] = l.Anchor
	}

	if l.HREFLang != "" {
		out["hreflang"] = l.HREFLang
	}

	if l.Media != "" {
		out["media"] = l.Media
	}

	if l.Title != "" {
		out["title"] = l.Title
	}

	if l.TitleStar != "" {
		out["title*"] = l.TitleStar
	}

	if l.Type != "" {
		out["type"] = l.Type
	}

	for _, extensionKey := range l.extensionKeys {
		out[extensionKey] = l.extensions[extensionKey]
	}

	return json.Marshal(out)

}
