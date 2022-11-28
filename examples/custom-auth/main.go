package main

import (
	"bufio"
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/service/clusters"
)

type CustomCredentials struct{}

func (c *CustomCredentials) Name() string {
	return "custom"
}

func (c *CustomCredentials) Configure(
	ctx context.Context, cfg *config.Config,
) (func(*http.Request) error, error) {
	return func(r *http.Request) error {
		token := askFor("Token:")
		r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		return nil
	}, nil
}

func main() {
	w := databricks.Must(databricks.NewWorkspaceClient(&databricks.Config{
		Host:        askFor("Host:"),
		Credentials: &CustomCredentials{},
	}))
	all, err := w.Clusters.ListAll(context.Background(), clusters.List{})
	if err != nil {
		panic(err)
	}
	for _, c := range all {
		println(c.ClusterName)
	}
}

func askFor(prompt string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stdout, prompt+" ")
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}
