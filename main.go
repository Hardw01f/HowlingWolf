package main

import (
	"context"
	"flag"
	"os"

	"github.com/google/subcommands"
	//"github.com/Snow-HardWolf/HowlingWolf/commands"
)

func main() {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&commands.testCmd{}, "")

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
