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
	if cfg.ResolveProfile && cfg.Host != "" && cfg.Profile == "" {
		err := cfg.fixHostIfNeeded()
		if err != nil {
			return err
		}
		profiles := []string{}
		for _, section := range iniFile.Sections() {
			values := section.KeysHash()
			host := values["host"]
			if host == "" {
				// if host is not set or is empty
				continue
			}
			canonical, err := canonicalHost(host)
			if err != nil {
				// we're fine with other corrupt profiles
				continue
			}
			if canonical != cfg.Host {
				continue
			}
			profiles = append(profiles, section.Name())
		}
		// in the real situations, we don't expect this to happen often (if not at all),
		// hence we don't trim the list
		if len(profiles) > 1 {
			return fmt.Errorf("%s match %s in %s",
				strings.Join(profiles, " and "),
				cfg.Host, configFile)
		}
		if len(profiles) == 1 {
			profile = profiles[0]
			logger.Debugf("%s config profile detected for %s",
				profile, cfg.Host)
		}
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
