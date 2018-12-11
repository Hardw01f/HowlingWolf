package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Server ServerConfig
}

type ServerConfig struct {
	Name string
}

func TomlLoader() {
	var config Config
	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		fmt.Println(err)
	}
}
