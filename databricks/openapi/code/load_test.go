package code

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombined(t *testing.T) {
	batch, err := NewFromFile("/private/tmp/processed-databricks-workspace-all.json")
	assert.NoError(t, err)

	for k, v := range batch.Packages["clusters"].services["Clusters"].methods {
		t.Logf("%s: %v", k, v.Shortcut())
	}

	// s := batch.Packages["permissions"].services["Permissions"].methods["GetObjectPermissions"].Shortcut()
	s := batch.Packages["clusters"].services["Clusters"].methods["create"].Shortcut()
	t.Logf("%v", s)

	w := batch.Packages["clusters"].services["Clusters"].methods["create"].Wait()
	b := w.Bind()
	t.Logf("%v", b)

	assert.Len(t, batch.Packages, 1)
}

func TestBasic(t *testing.T) {
	batch, err := NewFromFile("../testdata/spec.json")
	assert.NoError(t, err)

	assert.Len(t, batch.Packages, 1)
}
