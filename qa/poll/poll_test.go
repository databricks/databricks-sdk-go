package poll_test

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/qa/poll"
	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/stretchr/testify/assert"
)

func TestSimple(t *testing.T) {
	waiter := sql.WaitGetWarehouseRunning[int]{
		Poll: poll.Simple(sql.GetWarehouseResponse{
			Id: "test",
		}),
	}
	res, err := waiter.Get()
	assert.NoError(t, err)
	assert.Equal(t, "test", res.Id)
}
