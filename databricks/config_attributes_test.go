package databricks

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAttributesMaskedSignature(t *testing.T) {
	sig := ConfigAttributes.FieldNamesMask(&Config{
		AzureResourceID:    "irrelevant",
		DebugTruncateBytes: 1024,
		Host:               "irrelevant",
		GoogleCredentials:  "irrelevant",
		InsecureSkipVerify: true,
	})
	assert.Equal(t, "24:10a0602", sig)

	usedFields, err := strconv.ParseInt(strings.Split(sig, ":")[1], 16, 64)
	assert.NoError(t, err)

	flagHost := powerOfTwo(1).Int64()
	flagAccountID := powerOfTwo(2).Int64()
	flagToken := powerOfTwo(3).Int64()
	flagUsername := powerOfTwo(4).Int64()
	flagPassword := powerOfTwo(5).Int64()
	flagProfile := powerOfTwo(6).Int64()
	flagConfigFile := powerOfTwo(7).Int64()
	flagGoogleServiceAccount := powerOfTwo(8).Int64()
	flagGoogleCredentials := powerOfTwo(9).Int64()
	flagAzureWorkspaceResourceID := powerOfTwo(10).Int64()
	flagAzureUseMsi := powerOfTwo(11).Int64()
	flagAzureClientSecret := powerOfTwo(12).Int64()
	flagAzureClientID := powerOfTwo(13).Int64()
	flagAzureTenantID := powerOfTwo(14).Int64()
	flagAzureEnvironment := powerOfTwo(15).Int64()
	flagAuthType := powerOfTwo(16).Int64()
	flagSkipVerify := powerOfTwo(17).Int64()
	flagHttpTimeoutSeconds := powerOfTwo(18).Int64()
	flagDebugTruncateBytes := powerOfTwo(19).Int64()
	flagDebugHeaders := powerOfTwo(20).Int64()
	flagRateLimit := powerOfTwo(21).Int64()
	flagAzureLoginAppID := powerOfTwo(22).Int64()

	assert.True(t, usedFields&flagHost != 0)
	assert.False(t, usedFields&flagAccountID != 0)
	assert.False(t, usedFields&flagToken != 0)
	assert.False(t, usedFields&flagUsername != 0)
	assert.False(t, usedFields&flagPassword != 0)
	assert.False(t, usedFields&flagProfile != 0)
	assert.False(t, usedFields&flagConfigFile != 0)
	assert.False(t, usedFields&flagGoogleServiceAccount != 0)
	assert.True(t, usedFields&flagGoogleCredentials != 0)
	assert.True(t, usedFields&flagAzureWorkspaceResourceID != 0)
	assert.False(t, usedFields&flagAzureUseMsi != 0)
	assert.False(t, usedFields&flagAzureClientSecret != 0)
	assert.False(t, usedFields&flagAzureClientID != 0)
	assert.False(t, usedFields&flagAzureTenantID != 0)
	assert.False(t, usedFields&flagAzureEnvironment != 0)
	assert.False(t, usedFields&flagAuthType != 0)
	assert.True(t, usedFields&flagSkipVerify != 0)
	assert.False(t, usedFields&flagHttpTimeoutSeconds != 0)
	assert.True(t, usedFields&flagDebugTruncateBytes != 0)
	assert.False(t, usedFields&flagDebugHeaders != 0)
	assert.False(t, usedFields&flagRateLimit != 0)
	assert.False(t, usedFields&flagAzureLoginAppID != 0)
}
