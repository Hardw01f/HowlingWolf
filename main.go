package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"./commands"
	//"github.com/Snow-HardWolf/HowlingWolf/commands"
	"github.com/BurntSushi/toml"
	"github.com/google/subcommands"
)

type Config struct {
	Server ServerConfig
}

type ServerConfig struct {
	Name string
}

func main() {
	var config Config
	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		fmt.Println(err)
	}
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&commands.TestCmd{}, "")
	subcommands.Register(&commands.PsCmd{}, "")
	subcommands.Register(&commands.PortCmd{}, "")

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
