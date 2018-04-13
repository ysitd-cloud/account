package config

func NewConfigFromEnv() *Config {
	return &Config{
		Verbose:  false,
		Database: newDatabaseFromEnv(),
		Render:   newRenderFromEnv(),
	}
}
