package rfc8288

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHeaderParsing(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	links, err := ParseLinkHeaders("Context-Type: application/json\r\n")
	require.NoError(err)
	assert.Equal(0, len(links))

	links, err = ParseLinkHeaders("Context-Type: application/json\r\nLink: </foo>; rel=\"hello\"\r\nAccept: *\r\nLink: </bar1>; rel=\"item\", </bar2>; rel=\"collection\"")
	require.NoError(err)
	assert.Equal(3, len(links))
	assert.Equal("hello", links[0].Rel)
	assert.Equal("item", links[1].Rel)
	assert.Equal("collection", links[2].Rel)
}
