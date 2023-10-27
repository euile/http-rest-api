package apiserver

// Config ...
type Config struct {
	BindAddr string `toml:"bind_addr"`

	LogLevel string `toml:"log_level"`
}

// будет отдавать инифциализированный конфиг с дефолтными параметрами
// NewConfig ...
func NewConfiпg() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}
