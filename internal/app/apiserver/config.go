package apiserver

import "github.com/Pastafarian-maiden/rest-api-forum/internal/app/store"

//Config ...
type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLever string `toml: log_level`
	Store    *store.Config
}

//NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLever: "debug",
		Store:    store.NewConfig(),
	}
}
