package apiserver

type Config struct {
	BindAddr string `toml:"bindaddr"`
	LogLevel string `toml:"log_level"`
	//Store    *store.Config
}

// Newconfig
func NewConfig() *Config {
	return &Config{
		BindAddr: ":3030",
		LogLevel: "debug",
		//Store:    store.NewcConfig(),
	}
}
