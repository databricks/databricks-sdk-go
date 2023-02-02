package useragent

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchSemVer(t *testing.T) {
	assert.NoError(t, matchSemVer("1.2.3"))
	assert.Error(t, matchSemVer("1.2.3.4"))
	assert.Error(t, matchSemVer("1.2.3-foo"))
	assert.Error(t, matchSemVer("1.2"))
}

func TestMatchAlphanum(t *testing.T) {
	assert.NoError(t, matchAlphanum("foo"))
	assert.NoError(t, matchAlphanum("FOO"))
	assert.NoError(t, matchAlphanum("FOO123"))
	assert.NoError(t, matchAlphanum("foo_bar"))
	assert.NoError(t, matchAlphanum("foo-bar"))
	assert.Error(t, matchAlphanum("foo bar"))
	assert.Error(t, matchAlphanum("foo/bar"))
}

func TestMatchAlphanumOrSemVer(t *testing.T) {
	assert.NoError(t, matchAlphanumOrSemVer("foo"))
	assert.NoError(t, matchAlphanumOrSemVer("1.2.3"))
	assert.Error(t, matchAlphanumOrSemVer("foo/bar"))
	assert.Error(t, matchAlphanumOrSemVer("1/2/3"))
}
