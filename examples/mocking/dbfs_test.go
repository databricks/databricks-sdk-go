package sample

import (
	"context"
	"sample/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	_ "github.com/golang/mock/mockgen/model"
	"github.com/stretchr/testify/assert"

	"github.com/databricks/databricks-sdk-go/service/dbfs"
	"github.com/databricks/databricks-sdk-go/workspaces"
)

//go:generate mockgen -package=mocks -destination=mocks/dbfs.go github.com/databricks/databricks-sdk-go/service/dbfs DbfsService

func TestClusters(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.New MockDbfsService(ctrl)

	ctx := context.Background()
	m.EXPECT().
		Create(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_, _ any) (*dbfs.CreateResponse, error) {
			return &dbfs.CreateResponse{
				Handle: 123,
			}, nil
		})

	w := workspaces.New()
	w.Dbfs = m

	h, err := w.Dbfs.Create(ctx, dbfs.Create{
		Path: "/a/b/c",
	})
	assert.NoError(t, err)

	assert.Equal(t, int64(123), h.Handle)
}
