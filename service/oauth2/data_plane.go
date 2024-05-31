package oauth2

import (
	"strings"

	"golang.org/x/oauth2"
)

type DataPlaneHelper struct {
	infos  map[string]*DataPlaneInfo
	tokens map[string]*oauth2.Token
}

func (o *DataPlaneHelper) GetDataPlaneDetails(method string, params []string, refresh func(*DataPlaneInfo) (*oauth2.Token, error), infoGetter func() (*DataPlaneInfo, error)) (string, *oauth2.Token, error) {
	if o.infos == nil {
		o.infos = make(map[string]*DataPlaneInfo)
	}
	if o.tokens == nil {
		o.tokens = make(map[string]*oauth2.Token)
	}
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
	if !tokenOk || !token.Valid() {
		newToken, err := refresh(info)
		if err != nil {
			return "", nil, err
		}
		o.tokens[key] = newToken
		token = newToken
	}
	return info.EndpointUrl, token, nil
}

func (o *DataPlaneHelper) generateKey(method string, params []string) string {
	allElements := []string{method}
	allElements = append(allElements, params...)
	return strings.Join(allElements, "/")
}
