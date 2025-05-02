package config

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFileIDTokenSource(t *testing.T) {
	file, err := os.CreateTemp("", "subjectTokenFile")
	require.NoError(t, err)
	defer os.Remove(file.Name())

	err = os.WriteFile(file.Name(), []byte("token-1337"), 0777)
	require.NoError(t, err)

	source := &fileIDTokenSource{
		idTokenFilePath: file.Name(),
	}

	token, err := source.IDToken(context.Background(), "")
	require.NoError(t, err)
	require.Equal(t, "token-1337", token.Value)
}
