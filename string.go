package rfc8288

import (
	"fmt"
	"strings"
)

// String returns the Link in a format usable for HTTP Headers as defined by RFC8288
func (l Link) String() string {
	var result []string

	result = append(result, fmt.Sprintf(`<%s>`, l.HREF.String()))

	if l.Rel != "" {
		result = append(result, fmt.Sprintf(`rel="%s"`, l.Rel))
	}

	if l.HREFLang != "" {
		result = append(result, fmt.Sprintf(`hreflang="%s"`, l.HREFLang))
	}

	if l.Media != "" {
		result = append(result, fmt.Sprintf(`media="%s"`, l.Media))
	}

	if l.Title != "" {
		result = append(result, fmt.Sprintf(`title="%s"`, l.Title))
	}

	if l.TitleStar != "" {
		result = append(result, fmt.Sprintf(`title*="%s"`, l.TitleStar))
	}

	if l.Type != "" {
		result = append(result, fmt.Sprintf(`type="%s"`, l.Type))
	}

	for _, key := range l.extensionKeys {
		value := l.extensions[key]
		result = append(result, fmt.Sprintf(`%s="%s"`, key, value))
	}

	return strings.Join(result, "; ")
}
