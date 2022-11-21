package mocking

import (
	"context"
	"mocking/mocks"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	_ "github.com/golang/mock/mockgen/model"
	"github.com/stretchr/testify/assert"

	"github.com/databricks/databricks-sdk-go/databricks"
	"github.com/databricks/databricks-sdk-go/service/dbfs"
	"github.com/databricks/databricks-sdk-go/workspaces"
)

//go:generate go run github.com/golang/mock/mockgen@latest -package=mocks -destination=mocks/dbfs.go github.com/databricks/databricks-sdk-go/service/dbfs DbfsService

func TestDbfsHighLevelAPI(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDbfs := mocks.NewMockDbfsService(ctrl)

	ctx := context.Background()
	mockDbfs.EXPECT().Create(gomock.Any(), gomock.Eq(dbfs.Create{
		Path:      "/a/b/c",
		Overwrite: true,
	})).Return(&dbfs.CreateResponse{
		Handle: 123,
	}, nil)
	mockDbfs.EXPECT().AddBlock(gomock.Any(), gomock.Eq(dbfs.AddBlock{
		Handle: 123,
		Data:   "YWJj",
	}))
	mockDbfs.EXPECT().Close(gomock.Any(), gomock.Eq(dbfs.Close{
		Handle: 123,
	}))

	w := workspaces.New(databricks.NewMockConfig(nil))
	w.Dbfs.DbfsService = mockDbfs

	err := w.Dbfs.Overwrite(ctx, "/a/b/c", strings.NewReader("abc"))
	assert.NoError(t, err)
}
