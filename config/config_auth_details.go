package config

import (
	"fmt"
	"slices"
	"strings"
)

// AuthDetails contains the details of the authentication configuration.
type AuthDetails struct {
	AuthType      string            `json:"auth_type"`
	Username      string            `json:"username,omitempty"`
	Host          string            `json:"host,omitempty"`
	AccountID     string            `json:"account_id,omitempty"`
	Configuration AuthConfiguration `json:"configuration"`
}

func (a *AuthDetails) String() string {
	s := ""
	if a.Host != "" {
		s += fmt.Sprintf(`host=%s`, a.Host)
	}

	if a.AuthType != "" {
		s += fmt.Sprintf(", auth_type=%s", a.AuthType)
	}

	if a.Username != "" {
		s += fmt.Sprintf(", username=%s", a.Username)
	}

	if a.AccountID != "" {
		s += fmt.Sprintf(", account_id=%s", a.AccountID)
	}

	conf := a.Configuration.String()
	if conf != "" {
		s += fmt.Sprintf(`
Configuration:
%s`, conf)
	}

	return s
}

// AuthConfiguration is a map of attribute name to its configuration.
type AuthConfiguration map[string]*AttrConfig

func (c *AuthConfiguration) String() string {
	keys := make([]string, 0, len(*c))
	for k := range *c {
		keys = append(keys, k)
	}
	slices.Sort(keys)

	var conf []string
	for _, k := range keys {
		v := (*c)[k]
		conf = append(conf, fmt.Sprintf("- %s=%s", k, v))
	}
	return strings.Join(conf, "\n")
}

type AttrConfig struct {
	Value            string  `json:"value"`
	Source           *Source `json:"source"`
	AuthTypeMismatch bool    `json:"auth_type_mismatch"`
}

func (a *AttrConfig) String() string {
	s := fmt.Sprintf(`%s (from %s)`, a.Value, a.Source.String())
	if a.AuthTypeMismatch {
		s += " (not used by the current auth type)"
	}
	return s
}

func (c *Config) GetAuthDetails(showSensitive bool) AuthDetails {
	attrSet := make(map[string]*AttrConfig, 0)
	for _, a := range ConfigAttributes {
		if a.IsZero(c) {
			continue
		}

		attrSet[a.Name] = &AttrConfig{
			Value:            getValue(c, &a, showSensitive),
			Source:           c.getSource(&a),
			AuthTypeMismatch: a.HasAuthAttribute() && c.AuthType != "" && !slices.Contains(a.AuthTypes, c.AuthType),
		}
	}

	return AuthDetails{
		AuthType:      c.AuthType,
		Host:          c.Host,
		Configuration: attrSet,
	}
}

func (c *Config) SetAttrSource(attr *ConfigAttribute, source *Source) {
	if c.attrSource == nil {
		c.attrSource = make(map[string]*Source)
	}
	c.attrSource[attr.Name] = source
}

func (c *Config) getSource(a *ConfigAttribute) *Source {
	if c.attrSource == nil {
		return &Source{
			Type: SourceDynamicConfig,
		}
	}

	attrSource, ok := c.attrSource[a.Name]
	if !ok {
		return &Source{
			Type: SourceDynamicConfig,
		}
	}

	return attrSource
}

func getValue(cfg *Config, a *ConfigAttribute, showSensitive bool) string {
	if !showSensitive && a.Sensitive {
		return "********"
	}

	return a.GetString(cfg)
}
