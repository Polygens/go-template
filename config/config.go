package config

// Config contains all the configuration variables for this service
type Config struct {
	LogLevel string `mapstructure:"log_level"`
	Port     int
}
