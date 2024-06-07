package config

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/databricks/databricks-sdk-go/internal/env"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var cliDummy = &Config{Host: "https://abc.cloud.databricks.com/"}

const smallExecutable = `#!/bin/sh
echo hello world
`

const largeExecutable = `#!/bin/sh
cat <<EOF
{
	"access_token": "token",
	"token_type": "Bearer",
	"expiry": "2023-05-22T00:00:00.000000+00:00"
}
EOF
exit 0
`

const failFirstSucceedThereafter = `#!/bin/sh

# Check if a token file exists in the same directory as the script
if [! -f "$(dirname "$0")/.token_file" ]; then
    # If not, create the file and set the token
    echo "error: workspace auth not configured" >&2
    touch "$(dirname "$0")/.token_file"
    exit 1
fi

cat <<EOF
{
	"access_token": "token",
	"token_type": "Bearer",
	"expiry": "2023-05-22T00:00:00.000000+00:00"
}
EOF
exit 0
`

func writeDummyExecutable(t *testing.T, path, contents string, truncateSize int) {
	f, err := os.Create(filepath.Join(path, "databricks"))
	require.NoError(t, err)
	defer f.Close()
	err = os.Chmod(f.Name(), 0755)
	require.NoError(t, err)
	_, err = f.WriteString(contents)
	require.NoError(t, err)

	if truncateSize > 0 {
		err = f.Truncate(int64(truncateSize))
		require.NoError(t, err)
	}
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
	writeDummyExecutable(t, tmp, smallExecutable, 0)
	t.Setenv("PATH", tmp)

	aa := DatabricksCliCredentials{}
	_, err := aa.Configure(context.Background(), cliDummy)
	require.NoError(t, err)
}

func TestDatabricksCliCredentials_InstalledLegacyWithSymlink(t *testing.T) {
	// Create a dummy databricks executable.
	tmp1 := t.TempDir()
	tmp2 := t.TempDir()
	writeDummyExecutable(t, tmp1, smallExecutable, 0)
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
	writeDummyExecutable(t, tmp, largeExecutable, 1024 * 1024)
	t.Setenv("PATH", tmp+string(os.PathListSeparator)+os.Getenv("PATH"))

	aa := DatabricksCliCredentials{}
	_, err := aa.Configure(context.Background(), cliDummy)
	require.NoError(t, err)
}

func TestDatabricksCliCredentials_FallbackToAccountLevel(t *testing.T) {
	env.CleanupEnvironment(t)

	tmp := t.TempDir()
	writeDummyExecutable(t, tmp, failFirstSucceedThereafter, 0)
	t.Setenv("PATH", tmp+string(os.PathListSeparator)+os.Getenv("PATH"))

	aa := DatabricksCliCredentials{}
	_, err := aa.Configure(context.Background(), cliDummy)
	require.NoError(t, err)
}