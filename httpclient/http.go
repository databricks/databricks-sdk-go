package httpclient

import "net/http"

func (a *ApiClient) ToHttpClient() *http.Client {
	return &http.Client{
		Transport: a,
		Timeout:   a.config.HTTPTimeout,
	}
}
