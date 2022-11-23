package internal

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go"
)

const fullCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const hexCharset = "0123456789abcdef"

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

func GetEnvInt64OrSkipTest(t *testing.T, name string) int64 {
	v := GetEnvOrSkipTest(t, name)
	i, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		t.Skipf("`%s` is not int64: %s", v, err)
	}
	return i
}

// RandomEmail generates random email
func RandomEmail(prefix ...string) string {
	return fmt.Sprintf("%s@example.com", RandomName(
		append([]string{"sdk-go-"}, prefix...)...))
}

// RandomName gives random name with optional prefix. e.g. qa.RandomName("tf-")
func RandomName(prefix ...string) string {
	rand.Seed(time.Now().UnixNano())
	randLen := 12
	b := make([]byte, randLen)
	for i := range b {
		b[i] = fullCharset[rand.Intn(randLen)]
	}
	if len(prefix) > 0 {
		return fmt.Sprintf("%s%s", strings.Join(prefix, ""), b)
	}
	return string(b)
}

func RandomHex(prefix string, randLen int) string {
	rand.Seed(time.Now().UnixNano())

	b := make([]byte, randLen)
	for i := range b {
		b[i] = hexCharset[rand.Intn(randLen)%len(hexCharset)]
	}
	if len(prefix) > 0 {
		return fmt.Sprintf("%s%s", prefix, b)
	}
	return string(b)
}
