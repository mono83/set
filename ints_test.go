package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOfInts(t *testing.T) {
	s := OfInts(8, -42)
	if assert.NotNil(t, s) && assert.Len(t, s, 2) {
		assert.True(t, s[-42])
		assert.True(t, s[8])
		assert.False(t, s[11])
	}
}

func TestOfInts64(t *testing.T) {
	s := OfInts64(8, -42)
	if assert.NotNil(t, s) && assert.Len(t, s, 2) {
		assert.True(t, s[-42])
		assert.True(t, s[8])
		assert.False(t, s[11])
	}
}

func TestOfUints64(t *testing.T) {
	s := OfUints64(8, 42)
	if assert.NotNil(t, s) && assert.Len(t, s, 2) {
		assert.True(t, s[42])
		assert.True(t, s[8])
		assert.False(t, s[11])
	}
}
