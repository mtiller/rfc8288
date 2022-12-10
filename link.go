package rfc8288

import (
	"errors"
	"io"
	"net/url"
	"strings"
)

var (
	// ErrExtensionKeyIsReserved describes an attempt to call Link.Extend(k,v) with a reserved key name
	ErrExtensionKeyIsReserved = errors.New("rfc8288: the given extension key name is reserved please choose another name")

	// ReservedKeys holds the names of all the reserved key names that are not allowed to be used as extensions
	ReservedKeys = map[string]struct{}{
		"href":     {},
		"rel":      {},
		"hreflang": {},
		"media":    {},
		"title":    {},
		"title*":   {},
		"type":     {},
	}
)

// ParseLink attempts to parse a link string
func ParseLink(link string) (Link, error) {

	var (
		rs io.RuneScanner = strings.NewReader(link)
		s                 = scanner{runeScanner: rs}
		p                 = parser{scanner: s}
	)

	return p.parse()

}

// Link is an implementation of the structure defined by RFC8288 Web Linking
type Link struct {
	HREF      url.URL
	Rel       string
	HREFLang  string
	Media     string
	Title     string
	TitleStar string
	Type      string

	extensionKeys []string
	extensions    map[string]interface{}
}

// ExtensionKeys returns a slice of strings representing the names of extension keys for this Link struct in the order
// they were added
func (l Link) ExtensionKeys() []string {
	return l.extensionKeys
}

// Extension retrieves the value for an extension if present. A bool is also returned to signify whether the value was
// present upon retrieval
func (l *Link) Extension(key string) (interface{}, bool) {

	if l.extensions == nil {
		l.extensions = make(map[string]interface{})
	}

	val, ok := l.extensions[key]
	return val, ok

}

// Extend adds an extension to the Link. Only non-reserved extension keys are allowed.
// Setting the value to nil will remove the extension.
func (l *Link) Extend(key string, value interface{}) error {

	if _, reserved := ReservedKeys[strings.ToLower(key)]; reserved {
		return ErrExtensionKeyIsReserved
	}

	_, keyFound := l.Extension(key)
	if !keyFound {
		l.extensionKeys = append(l.extensionKeys, key)
	}

	if value != nil {
		l.extensions[key] = value
	} else {

		delete(l.extensions, key)

		for x := 0; x < len(l.extensionKeys); {

			if strings.EqualFold(key, l.extensionKeys[x]) {
				l.extensionKeys = append(l.extensionKeys[:x], l.extensionKeys[x+1:]...)
				break
			}

			x++

		}

	}

	return nil

}
