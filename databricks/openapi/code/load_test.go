package code

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasic(t *testing.T) {
	batch, err := NewFromFile("../testdata/spec.json")
	assert.NoError(t, err)

	assert.Len(t, batch.Packages, 1)
}

func TestBasic2(t *testing.T) {
	batch, err := NewFromFile("/tmp/processed-databricks-workspace-all.json")
	assert.NoError(t, err)

	r := batch.Packages["dbsql"].services["DataSources"].methods["listDataSources"].Response
	t.Log(r.CamelName())

	assert.Len(t, batch.Packages, 1)
}
