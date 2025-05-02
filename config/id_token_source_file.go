package config

import (
	"context"
	"errors"
	"os"

	"github.com/databricks/databricks-sdk-go/logger"
)

// fileIDTokenSource retrieves JWT Tokens from a file.
type fileIDTokenSource struct {
	idTokenFilePath string
}

// IDToken returns a JWT Token for the specified audience. It will return
// an error if not running in GitHub Actions.
func (g *fileIDTokenSource) IDToken(ctx context.Context, audience string) (*IDToken, error) {
	if g.idTokenFilePath == "" {
		logger.Debugf(ctx, "Missing IDTokenFilePath")
		return nil, errors.New("missing IDTokenFilePath")
	}

	data, err := os.ReadFile(g.idTokenFilePath)
	if err != nil {
		return nil, err
	}

	return &IDToken{Value: string(data)}, nil
}
