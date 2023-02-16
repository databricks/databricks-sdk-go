package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/databricks/databricks-sdk-go/logger"
	"gopkg.in/ini.v1"
)

type KnownConfigLoader struct{}

func (l KnownConfigLoader) Name() string {
	return "config-file"
}

func (l KnownConfigLoader) Configure(cfg *Config) error {
	// Skip loading config file if some authentication is already explicitly
	// configured directly in the config by a user.
	// See: https://github.com/databricks/databricks-sdk-go/issues/304
	if l.isAnyAuthConfigured(cfg) {
		return nil
	}
	configFile := cfg.ConfigFile
	if configFile == "" {
		configFile = "~/.databrickscfg"
	}
	if strings.HasPrefix(configFile, "~") {
		homedir, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("cannot find homedir: %w", err)
		}
		configFile = fmt.Sprintf("%s%s", homedir, configFile[1:])
	}
	_, err := os.Stat(configFile)
	if os.IsNotExist(err) {
		// early return for non-configured machines
		logger.Debugf("%s not found on current host", configFile)
		return nil
	}
	iniFile, err := ini.Load(configFile)
	if err != nil {
		return fmt.Errorf("cannot parse config file: %w", err)
	}
	// don't modify the config, so that debug appears clearly
	profile := cfg.Profile
	hasExplicitProfile := cfg.Profile != ""
	if !hasExplicitProfile {
		profile = "DEFAULT"
	}
	profileValues := iniFile.Section(profile)
	if len(profileValues.Keys()) == 0 {
		if !hasExplicitProfile {
			logger.Debugf("%s has no %s profile configured", configFile, profile)
			return nil
		}
		return fmt.Errorf("%s has no %s profile configured", configFile, profile)
	}
	logger.Debugf("Loading %s profile from %s", profile, configFile)
	err = ConfigAttributes.ResolveFromStringMap(cfg, profileValues.KeysHash())
	if err != nil {
		return fmt.Errorf("%s %s profile: %w", configFile, profile, err)
	}
	return nil
}

func (l KnownConfigLoader) isAnyAuthConfigured(cfg *Config) bool {
	for _, a := range ConfigAttributes {
		if a.Auth == "" {
			continue
		}
		if !a.IsZero(cfg) {
			return true
		}
	}
	return false
}
