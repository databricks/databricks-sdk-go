// Code generated from Databricks SDK for Go integration tests by openapi.roll.TestRegenerateExamples. DO NOT EDIT.

package settings_test

import (
	"context"
	"errors"
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/stretchr/testify/mock"
)

func TestSetting(t *testing.T) {
	w := mocks.NewMockWorkspaceClient(t)
	s := w.GetMockDefaultNamespaceAPI()
	s.EXPECT().UpdateDefaultNamespaceSetting(mock.Anything, mock.Anything).Return(nil, errors.New("error"))
	s.UpdateDefaultNamespaceSetting(context.Background(), settings.UpdateDefaultNamespaceSettingRequest{})

}
