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
	assert.NoError(t, matchAlphanum("foo"))
	assert.True(t, isAlphanum("foo"))

	assert.NoError(t, matchAlphanum("FOO"))
	assert.True(t, isAlphanum("FOO"))

	assert.NoError(t, matchAlphanum("FOO123"))
	assert.True(t, isAlphanum("FOO123"))

	assert.NoError(t, matchAlphanum("foo_bar"))
	assert.True(t, isAlphanum("foo_bar"))

	assert.NoError(t, matchAlphanum("foo-bar"))
	assert.True(t, isAlphanum("foo-bar"))

	assert.NoError(t, matchAlphanum("foo.bar"))
	assert.True(t, isAlphanum("foo.bar"))

	assert.Error(t, matchAlphanum("foo bar"))
	assert.False(t, isAlphanum("foo bar"))

	assert.Error(t, matchAlphanum("foo/bar"))
	assert.False(t, isAlphanum("foo/bar"))
}
