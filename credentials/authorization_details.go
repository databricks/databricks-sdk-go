package credentials

import (
	"net/url"
)

type AuthorizationDetails string

func (r AuthorizationDetails) EncodeValues(key string, v *url.Values) error {
	v.Set(key, string(r))
	return nil
}
