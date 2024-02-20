package main

import (
	"context"
	"crypto/tls"
	"log"
	"net/http"
	"os"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/databricks/databricks-sdk-go/service/compute"
)

func main() {
	if os.Getenv("HTTPS_PROXY") != "https://localhost:8443" && os.Getenv("HTTP_PROXY") != "http://localhost:8443" {
		log.Printf(`To run this example, first start the proxy server in the examples/http-proxy/proxy directory:

$ go run ../proxy [http|https]

Then, run this example setting the HTTP_PROXY or HTTPS_PROXY environment variable to the proxy server:

$ HTTPS_PROXY=https://localhost:8443 go run .
or
$ HTTP_PROXY=http://localhost:8443 go run .
depending on the proxy type you started.
`)
		os.Exit(1)
	}
	log.Printf("Constructing client...")
	logger.DefaultLogger = &logger.SimpleLogger{
		Level: logger.LevelDebug,
	}
	var tlsNextProto map[string]func(authority string, c *tls.Conn) http.RoundTripper
	if os.Getenv("HTTPS_PROXY") != "" {
		// Go's HTTP client only supports HTTP/1.1 proxies when using TLS. See
		// https://github.com/golang/go/issues/26479 for more information.
		tlsNextProto = make(map[string]func(authority string, c *tls.Conn) http.RoundTripper)
	}
	w := databricks.Must(databricks.NewWorkspaceClient(&databricks.Config{
		HTTPTransport: &http.Transport{
			TLSNextProto: tlsNextProto,
		},
	}))
	log.Printf("Listing clusters...")
	all, err := w.Clusters.ListAll(context.Background(), compute.ListClustersRequest{})
	if err != nil {
		panic(err)
	}
	for _, c := range all {
		println(c.ClusterName)
	}
}
