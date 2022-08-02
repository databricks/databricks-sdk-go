package databricks

import (
	"fmt"
	"log"
	"os"

	"github.com/mitchellh/go-homedir"
	"gopkg.in/ini.v1"
)

type KnownConfigLoader struct{}

func (l KnownConfigLoader) Name() string {
	return "config-file"
}

func (l KnownConfigLoader) Configure(cfg *Config) error {
	configFile := cfg.ConfigFile
	if configFile == "" {
		configFile = "~/.databrickscfg"
	}
	configFile, err := homedir.Expand(configFile)
	if err != nil {
		return fmt.Errorf("cannot find homedir: %w", err)
	}
	_, err = os.Stat(configFile)
	if os.IsNotExist(err) {
		// early return for non-configured machines
		log.Printf("[DEBUG] %s not found on current host", configFile)
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
			log.Printf("[DEBUG] %s has no %s profile configured", configFile, profile)
			return nil
		}
		return fmt.Errorf("%s has no %s profile configured", configFile, profile)
	}
	log.Printf("[INFO] Loading %s profile from %s", profile, configFile)
	err = ConfigAttributes.ResolveFromStringMap(cfg, profileValues.KeysHash())
	if err != nil {
		return fmt.Errorf("%s %s profile: %w", configFile, profile, err)
	}
	return nil
}
