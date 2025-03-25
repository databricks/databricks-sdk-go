package u2m

import (
	"fmt"
	"strings"
)

// WorkspaceOAuthArgument is an interface that provides the necessary information
// to authenticate using OAuth to a specific workspace.
type WorkspaceOAuthArgument interface {
	OAuthArgument

	// GetWorkspaceHost returns the host of the workspace to authenticate to.
	GetWorkspaceHost() string
}

// BasicWorkspaceOAuthArgument is a basic implementation of the WorkspaceOAuthArgument
// interface that links each host with exactly one OAuth token.
type BasicWorkspaceOAuthArgument struct {
	// host is the host of the workspace to authenticate to. This must start
	// with "https://" and must not have a trailing slash.
	host string
}

// NewBasicWorkspaceOAuthArgument creates a new BasicWorkspaceOAuthArgument.
func NewBasicWorkspaceOAuthArgument(host string) (BasicWorkspaceOAuthArgument, error) {
	if !strings.HasPrefix(host, "https://") {
		return BasicWorkspaceOAuthArgument{}, fmt.Errorf("host must start with 'https://': %s", host)
	}
	if strings.HasSuffix(host, "/") {
		return BasicWorkspaceOAuthArgument{}, fmt.Errorf("host must not have a trailing slash: %s", host)
	}
	return BasicWorkspaceOAuthArgument{host: host}, nil
}

// GetWorkspaceHost returns the host of the workspace to authenticate to.
func (a BasicWorkspaceOAuthArgument) GetWorkspaceHost() string {
	return a.host
}

// GetCacheKey returns a unique key for caching the OAuth token for the workspace.
// The key is in the format "<host>".
func (a BasicWorkspaceOAuthArgument) GetCacheKey() string {
	a.host = strings.TrimSuffix(a.host, "/")
	if !strings.HasPrefix(a.host, "http") {
		a.host = fmt.Sprintf("https://%s", a.host)
	}
	return a.host
}

var _ WorkspaceOAuthArgument = BasicWorkspaceOAuthArgument{}
