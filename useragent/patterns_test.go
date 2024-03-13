package useragent

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchSemVer(t *testing.T) {
	assert.NoError(t, matchSemVer("1.2.3"))
	assert.True(t, isSemVer("1.2.3"))

	assert.NoError(t, matchSemVer("0.0.0-dev+2e014739024a"))
	assert.True(t, isSemVer("0.0.0-dev+2e014739024a"))

	assert.Error(t, matchSemVer("1.2.3.4"))
	assert.False(t, isSemVer("1.2.3.4"))

	assert.Error(t, matchSemVer("1.2"))
	assert.False(t, isSemVer("1.2"))
}

func TestMatchAlphanum(t *testing.T) {
	assert.NoError(t, matchValidChars("foo"))
	assert.True(t, isValid("foo"))

	assert.NoError(t, matchValidChars("FOO"))
	assert.True(t, isValid("FOO"))

	assert.NoError(t, matchValidChars("FOO123"))
	assert.True(t, isValid("FOO123"))

	assert.NoError(t, matchValidChars("foo_bar"))
	assert.True(t, isValid("foo_bar"))

	assert.NoError(t, matchValidChars("foo-bar"))
	assert.True(t, isValid("foo-bar"))

	assert.NoError(t, matchValidChars("foo.bar"))
	assert.True(t, isValid("foo.bar"))

	assert.Error(t, matchValidChars("foo bar"))
	assert.False(t, isValid("foo bar"))

	assert.Error(t, matchValidChars("foo/bar"))
	assert.False(t, isValid("foo/bar"))
}
