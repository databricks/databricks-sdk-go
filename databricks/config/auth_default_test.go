package config_test

import (
	"context"
	"errors"
	"testing"

	"github.com/databricks/databricks-sdk-go/databricks"
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/databricks/env"
	"github.com/stretchr/testify/assert"
)

func TestErrCannotConfigureAuth(t *testing.T) {
	env.CleanupEnvironment(t)
	w := databricks.Must(databricks.NewWorkspaceClient())
	_, err := w.CurrentUser.Me(context.Background())
	assert.True(t, errors.Is(err, config.ErrCannotConfigureAuth))
}
