package credentials

import (
	"encoding/json"
	"errors"
	"net/url"
)

// Requited by query.Value. The AuthorizationDetails field (which is an array)
// must implement the EncodeValues method. It is not enought to have the
// array values implement the method.
type AuthorizationDetailsArray []AuthorizationDetails

func (r AuthorizationDetailsArray) EncodeValues(key string, v *url.Values) error {
	if key != "AuthorizationDetails" {
		return errors.New("invalid key: " + key)
	}
	jsonValue, err := json.Marshal(r)
	if err != nil {
		return err
	}
	v.Set("authorization_details", string(jsonValue))
	return nil
}

type AuthorizationDetails struct {
	Type       string   `json:"type" url:"type"`
	ObjectType string   `json:"object_type" url:"object_type"`
	ObjectPath string   `json:"object_path" url:"object_path"`
	Actions    []string `json:"actions" url:"actions"`
}
