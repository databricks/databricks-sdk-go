package client

import (
	"sync"
)

var (
	onceClient   sync.Once
	commonClient *DatabricksClient

	DefaultClient = &DatabricksClient{}
)

// ResetCommonEnvironmentClient resets test dummy
func ResetCommonEnvironmentClient() {
	commonClient = nil
	onceClient = sync.Once{}
}

// CommonEnvironmentClient configured once per run of application
func CommonEnvironmentClient() *DatabricksClient {
	if commonClient != nil {
		return commonClient
	}
	onceClient.Do(func() {
		commonClient = &DatabricksClient{}
		// err := commonClient.Authenticate(context.Background())
		// if err != nil {
		// 	panic(err)
		// }
	})
	return commonClient
}
