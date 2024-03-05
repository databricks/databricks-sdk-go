package config

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/databricks/databricks-sdk-go/logger"
	"gopkg.in/ini.v1"
)

// File represents the contents of a databrickscfg file.
type File struct {
	*ini.File

	path string
}

// Path returns the path of the loaded databrickscfg file.
func (f *File) Path() string {
	return f.path
}

func defaultConfigFilePath(path string) string {
	if path == "" {
		path = "~/.databrickscfg"
	}

	return path
}

// LoadFile loads the databrickscfg file at the specified path.
// The function loads ~/.databrickscfg if the specified path is an empty string.
// The function expands ~ to the user's home directory.
func LoadFile(path string) (*File, error) {
	path = defaultConfigFilePath(path)

	// Expand ~ to home directory.
	if strings.HasPrefix(path, "~") {
		homedir, err := os.UserHomeDir()
		if err != nil {
			return nil, fmt.Errorf("cannot find homedir: %w", err)
		}
		path = fmt.Sprintf("%s%s", homedir, path[1:])
	}

	iniFile, err := ini.LoadSources(ini.LoadOptions{
		// Passwords may contain '#'. Require a space before '#' for inline comments.
		// See https://github.com/databricks/databricks-sdk-go/issues/594
		SpaceBeforeInlineComment: true,
	}, path)
	if err != nil {
		return nil, err
	}

	return &File{iniFile, path}, err
}

var ConfigFile = configFileLoader{}

type configFileLoader struct{}

func (l configFileLoader) Name() string {
	return "config-file"
}

func (l configFileLoader) Configure(cfg *Config) error {
	// Skip loading config file if some authentication is already explicitly
	// configured directly in the config by a user.
	// See: https://github.com/databricks/databricks-sdk-go/issues/304
	if cfg.Profile == "" && (l.isAnyAuthConfigured(cfg) || cfg.Host != "" || cfg.AzureResourceID != "") {
		return nil
	}

	configFilePath := defaultConfigFilePath(cfg.ConfigFile)
	configFile, err := LoadFile(configFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			// early return for non-configured machines
			logger.Debugf(context.Background(), "%s not found on current host", configFilePath)
			return nil
		}
		return fmt.Errorf("cannot parse config file: %w", err)
	}
	// don't modify the config, so that debug appears clearly
	profile := cfg.Profile
	hasExplicitProfile := cfg.Profile != ""
	if !hasExplicitProfile {
		profile = "DEFAULT"
	}
	profileValues := configFile.Section(profile)
	if len(profileValues.Keys()) == 0 {
		if !hasExplicitProfile {
			logger.Debugf(context.Background(), "%s has no %s profile configured", configFile.Path(), profile)
			return nil
		}
		return fmt.Errorf("%s has no %s profile configured", configFile.Path(), profile)
	}
	logger.Debugf(context.Background(), "Loading %s profile from %s", profile, configFile.Path())
	err = ConfigAttributes.ResolveFromStringMapWithSource(cfg, profileValues.KeysHash(), Source{Type: SourceFile, Name: configFile.Path()})
	if err != nil {
		return fmt.Errorf("%s %s profile: %w", configFile.Path(), profile, err)
	}
	return nil
}

func (l configFileLoader) isAnyAuthConfigured(cfg *Config) bool {
	for _, a := range ConfigAttributes {
		if !a.HasAuthAttribute() {
			continue
		}
		if !a.IsZero(cfg) {
			return true
		}
	}
	return false
}
