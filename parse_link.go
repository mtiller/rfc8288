package rfc8288

import (
	"io"
	"strings"
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
