package commands

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
)

type testCmd struct {
	sentence string
}

func (*testCmd) Name() string     { return "test" }
func (*testCmd) Synopsis() string { return "test args to stdout." }
func (*testCmd) Usage() string {
	return `test [-capitalize] <some text>:
  Print a famous words
`
}

func (p *testCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.sentence, "Test", "", "testtest")
}

func (p *testCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	fmt.Println("GANDOM")
	return subcommands.ExitSuccess
}
