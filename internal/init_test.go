package internal

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/databricks"
	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/retries"
	"github.com/databricks/databricks-sdk-go/service/clusters"
	"github.com/stretchr/testify/assert"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
	databricks.WithProduct("integration-tests", databricks.Version())
}

// GetEnvOrSkipTest proceeds with test only with that env variable
func GetEnvOrSkipTest(t *testing.T, name string) string {
	value := os.Getenv(name)
	if value == "" {
		t.Skipf("Environment variable %s is missing", name)
	}
	return value
}

// RandomEmail generates random email
func RandomEmail(prefix ...string) string {
	return fmt.Sprintf("%s@example.com", RandomName(
		append([]string{"sdk-go-"}, prefix...)...))
}

// RandomName gives random name with optional prefix. e.g. qa.RandomName("tf-")
func RandomName(prefix ...string) string {
	randLen := 12
	b := make([]byte, randLen)
	for i := range b {
		b[i] = charset[rand.Intn(randLen)]
	}
	if len(prefix) > 0 {
		return fmt.Sprintf("%s%s", strings.Join(prefix, ""), b)
	}
	return string(b)
}

func startDefaultTestCluster(t *testing.T, ctx context.Context, apiClient *client.DatabricksClient) {
	clusterService := clusters.New(apiClient)
	clusterId := GetEnvOrSkipTest(t, "TEST_DEFAULT_CLUSTER_ID")
	err := retries.Wait(ctx, 10*time.Minute, func() *retries.Err {
		clusterDetails, err := clusterService.GetCluster(ctx, clusters.GetClusterRequest{
			ClusterId: clusterId,
		})
		if err != nil {
			return retries.Halt(err)
		}
		if clusterDetails.State == clusters.GetClusterResponseStateRunning {
			return nil
		}
		if clusterDetails.State == clusters.GetClusterResponseStateTerminated {
			err = clusterService.StartCluster(ctx, clusters.StartClusterRequest{
				ClusterId: clusterId,
			})
			if err != nil {
				return retries.Halt(err)
			}
		}
		return retries.Continue(fmt.Errorf("%s", clusterDetails.State))
	})
	assert.NoError(t, err)
}
