package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/Pastafarian-maiden/rest-api-forum/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
	//путь к файлу, который является конфигом, чтобы мы могли запускать его флагом вместе с запуском бинарника
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config) //записать значения с конфиг файла в Config
	if err != nil {
		log.Fatal(err)
	}
	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
