package compute

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrimLeadingWhitespace(t *testing.T) {
	assert.Equal(t, "foo\nbar\n", TrimLeadingWhitespace(`
	
    	    foo
            bar
	
	`))
}
