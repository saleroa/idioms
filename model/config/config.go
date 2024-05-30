package config

type Config struct {
	Database *Database `mapstructure:"database" yaml:"database"`
	Logger   *Logger   `mapstructure:"logger" yaml:"logger"`
	Server   *Server   `mapstructure:"server" yaml:"server"`
	Salt     string    `mapstructure:"salt" yaml:"salt"`
}
