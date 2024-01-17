package store

type Config struct {
	DataBaseURL string `toml:"database_url"`
}

func NewcConfig() *Config {
	return &Config{}
}
