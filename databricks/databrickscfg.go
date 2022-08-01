package databricks

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mitchellh/go-homedir"
	"gopkg.in/ini.v1"
)

type DatabricksCliCredentials struct {
	// Location of the Databricks CLI credentials file, that is created
	// by `databricks configure --token` command. By default, it is located
	// in ~/.databrickscfg.
	ConfigFile string `name:"config_file" env:"DATABRICKS_CONFIG_FILE"`

	// Connection profile specified within ~/.databrickscfg.
	Profile string `name:"profile" env:"DATABRICKS_CONFIG_PROFILE" auth:"config profile"`
}

func (c DatabricksCliCredentials) Name() string {
	return "databricks-cli"
}

func (c DatabricksCliCredentials) Configure(ctx context.Context, cfg *Config) (func(*http.Request) error, error) {
	configFile := c.ConfigFile
	if configFile == "" {
		configFile = "~/.databrickscfg"
	}
	configFile, err := homedir.Expand(configFile)
	if err != nil {
		return nil, fmt.Errorf("cannot find homedir: %w", err)
	}
	_, err = os.Stat(configFile)
	if os.IsNotExist(err) {
		// early return for non-configured machines
		log.Printf("[DEBUG] %s not found on current host", configFile)
		return nil, nil
	}
	iniFile, err := ini.Load(configFile)
	if err != nil {
		return nil, fmt.Errorf("cannot parse config file: %w", err)
	}
	if c.Profile == "" {
		log.Printf("[INFO] Using DEFAULT profile from %s", configFile)
		c.Profile = "DEFAULT"
	}
	profileValues := iniFile.Section(c.Profile)
	if len(profileValues.Keys()) == 0 {
		// here we meet a heavy user of Databricks CLI
		return nil, fmt.Errorf("%s has no %s profile configured", configFile, c.Profile)
	}
	visitor, err := DefaultCredentials{
		Explicit: profileValues.KeysHash(),
		skip:     c.Name(),
	}.Configure(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("%s %s profile: %w", configFile, c.Profile, err)
	}
	return visitor, nil
}
