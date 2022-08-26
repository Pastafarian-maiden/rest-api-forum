package apiserver

import "github.com/Pastafarian-maiden/rest-api-forum/internal/app/store"

type Config struct {
	BinAddr  string `toml:"bind_addr"` //TOML is a file format for configuration files.
	LogLevel string `toml:"log_level"`
	Store    *store.Config
}

func NewConfig() *Config {
	return &Config{
		BinAddr:  ":8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
