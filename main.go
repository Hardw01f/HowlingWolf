package main

import (
	"context"
	"flag"
	"os"

	//"./commands"
	"github.com/Snow-HardWolf/HowlingWolf/commands"
	"github.com/google/subcommands"
)

func main() {
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
