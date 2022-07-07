package implicit

import (
	"context"
	"net/http"

	"github.com/databricks/sdk-go/core/auth/azure"
	"github.com/databricks/sdk-go/core/auth/gcp"
	"github.com/databricks/sdk-go/core/auth/native"
	"github.com/databricks/sdk-go/core/client"
)

type Credentials interface {
	Name() string
	Configure(context.Context, *client.ApiClient) (func(*http.Request) error, error)
}

var Methods = []Credentials{
	native.PatCredentials{},
	azure.ClientSecretCredentials{},
	gcp.ImpersonateWorkspaceCredentials{},
	gcp.ImpersonateAccountsCredentials{},
	gcp.GoogleCredentials{},
	azure.AzureCliCredentials{},
}

var ConfigAttributes = combineAttrs(Methods)

func combineAttrs(cc []Credentials) (res []ConfigAttribute) {
	attrs := map[string]ConfigAttribute{}
	for _, m := range cc {
		method := m.Name()
		for _, a := range attributesOf(m, method) {
			existing, ok := attrs[a.Name]
			if ok {
				existing.Credentials = append(existing.Credentials, method)
				attrs[a.Name] = existing
				continue
			}
			attrs[a.Name] = a
		}
	}
	for _, v := range attrs {
		res = append(res, v)
	}
	return
}

// func (c *ApiClient) niceAuthError(message string) error {
// 	info := ""
// 	// TODO: move to implicit
// 	if len(c.configAttributesUsed) > 0 {
// 		envs := []string{}
// 		attrs := []string{}
// 		usedAsEnv := map[string]bool{}
// 		for _, attr := range ClientAttributes() {
// 			if len(attr.EnvVars) == 0 {
// 				continue
// 			}
// 			for _, envVar := range attr.EnvVars {
// 				value := os.Getenv(envVar)
// 				if value == "" {
// 					continue
// 				}
// 				usedAsEnv[attr.Name] = true
// 				envs = append(envs, envVar)
// 			}
// 		}
// 		for _, attr := range c.configAttributesUsed {
// 			if usedAsEnv[attr] {
// 				continue
// 			}
// 			attrs = append(attrs, attr)
// 		}
// 		infos := []string{}
// 		if len(attrs) > 0 {
// 			infos = append(infos, fmt.Sprintf("Attributes used: %s", strings.Join(attrs, ", ")))
// 		}
// 		if len(envs) > 0 {
// 			infos = append(infos, fmt.Sprintf("Environment variables used: %s", strings.Join(envs, ", ")))
// 		}
// 		info = ". " + strings.Join(infos, ". ")
// 	}
// 	info = strings.TrimSuffix(info, ".")
// 	message = strings.TrimSuffix(message, ".")
// 	return fmt.Errorf("%s%s", message, info)
// }
