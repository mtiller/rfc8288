package rfc8288

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtends(t *testing.T) {
	assert := assert.New(t)
	// should make the value accessible via Extension(key)"

	// given
	l := Link{}
	key := "extension"
	value := "value"

	// when
	l.Extend(key, value)
	result, ok := l.Extension(key)

	// then
	assert.True(ok)
	assert.Equal(value, result)

	// should delete the extension if assigned a nil value"

	// given
	l = Link{}
	key = "extension"
	value = "value"

	// when
	l.Extend(key, value)
	l.Extend(key, nil)
	result, ok = l.Extension(key)

	// then
	assert.False(ok)
	assert.Nil(result)

	// should return ErrExtensionKeyIsReserved if attempting to Extend with reserved key"

	// given
	l = Link{}

	// expect
	for reservedKey := range ReservedKeys {
		err := l.Extend(reservedKey, "any")
		assert.Equal(ErrExtensionKeyIsReserved, err)
	}
}
