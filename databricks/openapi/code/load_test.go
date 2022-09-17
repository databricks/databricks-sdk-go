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

// This test is used for debugging purposes
func TestBasicDebug(t *testing.T) {
	//t.SkipNow()
	batch, err := NewFromFile("/tmp/processed-databricks-workspace-all.json")
	assert.NoError(t, err)

	// m := batch.Packages["commands"].services["CommandExecution"].methods["cancel"]
	m := batch.Packages["jobs"].services["Jobs"].methods["repairRun"]
	w := m.Wait()
	x := w.Binding()
	t.Log(x)
	y := x[0].Bind.Of.CamelName()
	t.Log(y)
	//m := batch.Packages["unitycatalog"].services["Catalogs"].methods["create"]
	t.Log(m)

	assert.Len(t, batch.Packages, 1)
}
