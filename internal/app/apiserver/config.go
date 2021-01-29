package apiserver

// Config ...
type Config struct {
	BindAddr    string `toml:"bind_addr"`
	LogLevel    string `toml:"log_level"`
	Database    database
	// DatabaseURL string `toml:"database_ur"`
	SessionKey  string `toml:"session_key"`
}

type database struct {
	Server   string `toml:"server"`
	Port     string `toml:"port"`
	Database string `toml:"database"`
	User     string `toml:"user"`
	Password string `toml:"password"`
}

// NewConfig ...
func NewConfig() *Config {
	// default value
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}
