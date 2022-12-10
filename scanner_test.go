package rfc8288

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type tokenLiteral struct {
	Token   token
	Literal string
}

type scannerTestEntry struct {
	name   string
	input  string
	tokens []tokenLiteral
}

func scannerEntry(name string, input string, tokens []tokenLiteral) scannerTestEntry {
	return scannerTestEntry{
		name:   name,
		input:  input,
		tokens: tokens,
	}
}

func TestScanner(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	entries := []scannerTestEntry{
		scannerEntry(
			"example 1",
			`<about:blank>; rel="prev"; title*="title"; media="media"; custom*="custom"`,
			[]tokenLiteral{
				{Token: LT, Literal: "<"},
				{Token: WORD, Literal: "about:blank"},
				{Token: GT, Literal: ">"},
				{Token: SEMICOLON, Literal: ";"},
				{Token: WS, Literal: " "},
				{Token: REL, Literal: "rel"},
				{Token: EQ, Literal: "="},
				{Token: QUOTE, Literal: `"`},
				{Token: WORD, Literal: "prev"},
				{Token: QUOTE, Literal: `"`},
				{Token: SEMICOLON, Literal: ";"},
				{Token: WS, Literal: " "},
				{Token: TITLE, Literal: "title"},
				{Token: STAR, Literal: "*"},
				{Token: EQ, Literal: "="},
				{Token: QUOTE, Literal: `"`},
				{Token: WORD, Literal: "title"},
				{Token: QUOTE, Literal: `"`},
				{Token: SEMICOLON, Literal: ";"},
				{Token: WS, Literal: " "},
				{Token: MEDIA, Literal: "media"},
				{Token: EQ, Literal: "="},
				{Token: QUOTE, Literal: `"`},
				{Token: WORD, Literal: "media"},
				{Token: QUOTE, Literal: `"`},
				{Token: SEMICOLON, Literal: ";"},
				{Token: WS, Literal: " "},
				{Token: WORD, Literal: "custom"},
				{Token: STAR, Literal: "*"},
				{Token: EQ, Literal: "="},
				{Token: QUOTE, Literal: `"`},
				{Token: WORD, Literal: "custom"},
				{Token: QUOTE, Literal: `"`},
				{Token: EOF, Literal: ""},
			},
		),
		scannerEntry(
			"example two. it's ok that this is an invalid link. lexer don't care",
			"<https://www.google.com> asdf",
			[]tokenLiteral{
				{Token: LT, Literal: "<"},
				{Token: WORD, Literal: "https://www.google.com"},
				{Token: GT, Literal: ">"},
				{Token: WS, Literal: " "},
				{Token: WORD, Literal: "asdf"},
				{Token: EOF, Literal: ""},
			},
		),
		scannerEntry(
			"Edge Case: Ends with whitespace",
			" ",
			[]tokenLiteral{
				{Token: WS, Literal: " "},
				{Token: EOF, Literal: ""},
			},
		),
	}
	for _, entry := range entries {
		in := entry.input
		out := entry.tokens

		// given
		r := strings.NewReader(in)
		s := scanner{runeScanner: r}

		x := 0
		for {
			// assert that we haven't scanned more than we expect to
			assert.Greater(len(out), x, fmt.Sprintf("check that we haven't scanned more than we expect in case %s", entry.name))

			// when
			token, literal, err := s.Scan()

			// then
			require.NoError(err)
			assert.Equal(out[x].Token, token)
			assert.Equal(out[x].Literal, literal)

			x++

			if token == EOF {
				break
			}
		}

		assert.Equal(len(out), x)
	}
}
