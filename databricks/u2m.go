package databricks

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/oauth2/authhandler"

	"github.com/databricks/databricks-sdk-go/databricks/internal"
	"github.com/databricks/databricks-sdk-go/databricks/oauth"
)


type DatabricksOAuthU2M struct {
}

func (c DatabricksOAuthU2M) Name() string {
	return "oauth-u2m"
}

func (c DatabricksOAuthU2M) Configure(ctx context.Context, cfg *Config) (func(*http.Request) error, error) {
	if cfg.ClientID == "" || cfg.Host == "" {
		return nil, nil
	}
	refreshCtx := context.Background()

	feedback := make(chan [2]string)
	srv := &http.Server{
		Addr: "localhost:8020",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			feedback <- [2]string{r.FormValue("code"), r.FormValue("state")}
			w.WriteHeader(200)
			w.Write([]byte("done"))
		}),
	}
	go srv.ListenAndServe()
	config, err := cfg.OAuthConfig()
	if err != nil {
		return nil, err
	}
	state, pkce := oauth.StateAndPKCE()
	ts := authhandler.TokenSourceWithPKCE(refreshCtx, config, state, func(authCodeURL string) (code string, state string, err error) {
		fmt.Println(authCodeURL)
		ctx, cancel := context.WithTimeout(refreshCtx, 30*time.Second)
		defer cancel()
		select {
		case <-ctx.Done():
			return "", "", fmt.Errorf("timed out")
		case success := <-feedback:
			srv.Close()
			return success[0], success[1], nil
		}
	}, pkce)
	token, err := ts.Token()
	if err != nil {
		return nil, fmt.Errorf("token: %w", err)
	}
	fmt.Printf("%v", token) // TODO: wait 5 minutes between requests to verify it works
	return internal.RefreshableVisitor(ts), nil
}
