package databricks

// Default settings
const (
	defaultDebugTruncateBytes = 96
	defaultRateLimitPerSecond = 15
	defaultHTTPTimeoutSeconds = 60
)

type TransportDefaults struct{}

func (l TransportDefaults) Name() string {
	return "transport-defaults"
}

func (l TransportDefaults) Configure(cfg *Config) error {
	if cfg.DebugTruncateBytes == 0 {
		cfg.DebugTruncateBytes = defaultDebugTruncateBytes
	}
	if cfg.HTTPTimeoutSeconds == 0 {
		cfg.HTTPTimeoutSeconds = defaultHTTPTimeoutSeconds
	}
	if cfg.RateLimitPerSecond == 0 {
		cfg.RateLimitPerSecond = defaultRateLimitPerSecond
	}
	return nil
}
