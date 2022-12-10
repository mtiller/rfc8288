package rfc8288

import (
	"bufio"
	"fmt"
	"net/textproto"
	"strings"
)

func LinkHeader(links ...Link) string {
	values := []string{}
	for _, link := range links {
		values = append(values, link.String())
	}
	return fmt.Sprintf("Link: %s", strings.Join(values, ", "))
}

func ParseLinkHeaders(v string) ([]Link, error) {
	reader := bufio.NewReader(strings.NewReader(v + "\r\n\r\n"))
	tp := textproto.NewReader(reader)
	mimeHeader, err := tp.ReadMIMEHeader()
	if err != nil {
		return []Link{}, fmt.Errorf("Error reading MIME header: %s", err.Error())
	}
	ret := []Link{}
	for key, val := range mimeHeader {
		if key == "Link" {
			for _, hval := range val {
				values := strings.Split(hval, ",")
				for _, lval := range values {
					link, err := ParseLink(lval)
					if err != nil {
						return []Link{}, err
					}
					ret = append(ret, link)
				}
			}
		}
	}
	return ret, nil
}
