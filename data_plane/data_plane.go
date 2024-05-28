package dataplane

import (
	"strings"

	dp "github.com/databricks/databricks-sdk-go/service/oauth2"
	"golang.org/x/oauth2"
)

type DataPlaneTokenCache struct {
	infos  map[string]*dp.DataPlaneInfo
	tokens map[string]*oauth2.Token
}

func (o *DataPlaneTokenCache) GetDataPlane(method string, params []string, refresh func(*dp.DataPlaneInfo) (*oauth2.Token, error), infoGetter func() (*dp.DataPlaneInfo, error)) (string, *oauth2.Token, error) {
	key := o.generateKey(method, params)
	info, infoOk := o.infos[key]
	token, tokenOk := o.tokens[key]
	if infoOk && tokenOk && token.Valid() {
		return info.EndpointUrl, token, nil
	}
	if !infoOk {
		newInfo, err := infoGetter()
		if err != nil {
			return "", nil, err
		}
		o.infos[key] = newInfo
		info = newInfo
	}
	newToken, err := refresh(info)
	if err != nil {
		return "", nil, err
	}
	o.tokens[method] = newToken
	return info.EndpointUrl, newToken, nil
}

func (o *DataPlaneTokenCache) generateKey(method string, params []string) string {
	allElements := append(params, method)
	return strings.Join(allElements, "/")
}
