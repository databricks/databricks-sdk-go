package config_test

import (
	"context"
	"errors"
	"github.com/xuxiaoshuo/databricks-sdk-go"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuxiaoshuo/databricks-sdk-go/config"
	"github.com/xuxiaoshuo/databricks-sdk-go/internal/env"
)

func TestErrCannotConfigureAuth(t *testing.T) {
	env.CleanupEnvironment(t)
	w := databricks.Must(databricks.NewWorkspaceClient())
	_, err := w.CurrentUser.Me(context.Background())
	assert.True(t, errors.Is(err, config.ErrCannotConfigureAuth))
}
