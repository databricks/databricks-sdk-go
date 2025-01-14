package config

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/databricks/databricks-sdk-go/databricks/env"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var cliDummy = &Config{Host: "https://abc.cloud.databricks.com/"}

func writeSmallDummyExecutable(t *testing.T, path string) {
	f, err := os.Create(filepath.Join(path, "databricks"))
	require.NoError(t, err)
	defer f.Close()
	err = os.Chmod(f.Name(), 0755)
	require.NoError(t, err)
	_, err = f.WriteString("#!/bin/sh\necho hello world\n")
	require.NoError(t, err)
}

func writeLargeDummyExecutable(t *testing.T, path string) {
	f, err := os.Create(filepath.Join(path, "databricks"))
	require.NoError(t, err)
	defer f.Close()
	err = os.Chmod(f.Name(), 0755)
	require.NoError(t, err)
	_, err = f.WriteString("#!/bin/sh\n")
	require.NoError(t, err)

	f.WriteString(`
cat <<EOF
{
"access_token": "token",
"token_type": "Bearer",
"expiry": "2023-05-22T00:00:00.000000+00:00"
}
EOF
`)
	require.NoError(t, err)
	_, err = f.WriteString("exit 0\n")
	require.NoError(t, err)

	err = f.Truncate(1024 * 1024)
	require.NoError(t, err)
}

func TestDatabricksCliCredentials_SkipAzure(t *testing.T) {
	aa := DatabricksCliCredentials{}
	x, err := aa.Configure(context.Background(), &Config{Host: "https://adb-xyz.c.azuredatabricks.net/"})
	assert.Nil(t, x)
	assert.NoError(t, err)
}

func TestDatabricksCliCredentials_SkipGcp(t *testing.T) {
	aa := DatabricksCliCredentials{}
	x, err := aa.Configure(context.Background(), &Config{Host: "https://123.4.gcp.databricks.com/"})
	assert.Nil(t, x)
	assert.NoError(t, err)
}

func TestDatabricksCliCredentials_NotInstalled(t *testing.T) {
	t.Setenv("PATH", "whatever")
	aa := DatabricksCliCredentials{}
	_, err := aa.Configure(context.Background(), cliDummy)
	require.NoError(t, err)
}

func TestDatabricksCliCredentials_InstalledLegacy(t *testing.T) {
	// Create a dummy databricks executable.
	tmp := t.TempDir()
	writeSmallDummyExecutable(t, tmp)
	t.Setenv("PATH", tmp)

	aa := DatabricksCliCredentials{}
	_, err := aa.Configure(context.Background(), cliDummy)
	require.NoError(t, err)
}

func TestDatabricksCliCredentials_InstalledLegacyWithSymlink(t *testing.T) {
	// Create a dummy databricks executable.
	tmp1 := t.TempDir()
	tmp2 := t.TempDir()
	writeSmallDummyExecutable(t, tmp1)
	os.Symlink(filepath.Join(tmp1, "databricks"), filepath.Join(tmp2, "databricks"))
	t.Setenv("PATH", tmp2+string(os.PathListSeparator)+os.Getenv("PATH"))

	aa := DatabricksCliCredentials{}
	_, err := aa.Configure(context.Background(), cliDummy)
	require.NoError(t, err)
}

func TestDatabricksCliCredentials_InstalledNew(t *testing.T) {
	env.CleanupEnvironment(t)

	// Create a dummy databricks executable.
	tmp := t.TempDir()
	writeLargeDummyExecutable(t, tmp)
	t.Setenv("PATH", tmp+string(os.PathListSeparator)+os.Getenv("PATH"))

	aa := DatabricksCliCredentials{}
	_, err := aa.Configure(context.Background(), cliDummy)
	require.NoError(t, err)
}
