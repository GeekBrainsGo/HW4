package webserver

// Config ...
type Config struct {
	BindAddr                 string `toml:"bind_addr"`
	LogLevel                 string `toml:"log_level"`
	DatabaseConnectionString string `toml:"dsn_url"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8888",
		LogLevel: "debug",
	}
}
