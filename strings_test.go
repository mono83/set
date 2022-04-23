package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOfStrings(t *testing.T) {
	s := OfStrings("foo", "bar")
	if assert.NotNil(t, s) && assert.Len(t, s, 2) {
		assert.True(t, s["foo"])
		assert.True(t, s["bar"])
		assert.False(t, s["baz"])
	}
}
