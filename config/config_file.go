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
	configFile, err := l.resolveConfigFile(cfg)
	if os.IsNotExist(err) {
		// early return for non-configured machines
		logger.Debugf("%s not found on current host", configFile)
		return nil
	}
	if err != nil {
		return fmt.Errorf(".databrickscfg: %w", err)
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

func (l KnownConfigLoader) resolveConfigFile(cfg *Config) (string, error) {
	configFile := cfg.ConfigFile
	if configFile == "" {
		configFile = "~/.databrickscfg"
	}
	if strings.HasPrefix(configFile, "~") {
		homedir, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("cannot find homedir: %w", err)
		}
		configFile = fmt.Sprintf("%s%s", homedir, configFile[1:])
	}
	_, err := os.Stat(configFile)
	if err != nil {
		return "", err
	}
	return configFile, nil
}

// ImplicitProfileLoader is an opt-in loader, which triggers if you specify
// `host`, but no other credentials either through direct configuration or
// through environment variables, Databricks SDK for Go will try picking up
// profile with the matching host from ~/.databrickscfg. This allows keeping
// the hostname checked in to version control, but have ability to pick up
// different credentials either from local development machine or production
// server.
type ImplicitProfileLoader struct {
	KnownConfigLoader
}

func (l ImplicitProfileLoader) Name() string {
	return "implicit-profile"
}

func (l ImplicitProfileLoader) resolveSection(
	cfg *Config, iniFile *ini.File) (*ini.Section, error) {
	var candidates []*ini.Section
	for _, section := range iniFile.Sections() {
		hash := section.KeysHash()
		host, ok := hash["host"]
		if !ok {
			// if host is not set
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
		candidates = append(candidates, section)
	}
	// in the real situations, we don't expect this to happen often
	// (if not at all), hence we don't trim the list
	if len(candidates) > 1 {
		var profiles []string
		for _, v := range candidates {
			profiles = append(profiles, v.Name())
		}
		return nil, fmt.Errorf("%s match %s in %s",
			strings.Join(profiles, " and "), cfg.Host, cfg.ConfigFile)
	}
	return candidates[0], nil
}

func (l ImplicitProfileLoader) Configure(cfg *Config) error {
	if cfg.Host == "" || cfg.Profile != "" {
		// no host to match or profile already configured
		return nil
	}
	configFile, err := l.resolveConfigFile(cfg)
	if os.IsNotExist(err) {
		// early return for non-configured machines
		logger.Debugf("%s not found on current host", configFile)
		return nil
	}
	if err != nil {
		return fmt.Errorf(".databrickscfg: %w", err)
	}
	iniFile, err := ini.Load(configFile)
	if err != nil {
		return fmt.Errorf("cannot parse config file: %w", err)
	}
	err = cfg.fixHostIfNeeded()
	if err != nil {
		return err
	}
	section, err := l.resolveSection(cfg, iniFile)
	if err != nil {
		return err
	}
	profile := section.Name()
	logger.Debugf("Loading %s profile from %s", profile, configFile)
	err = ConfigAttributes.ResolveFromStringMap(cfg, section.KeysHash())
	if err != nil {
		return fmt.Errorf("%s %s profile: %w", configFile, profile, err)
	}
	return nil
}
