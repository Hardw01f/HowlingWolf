package commands

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/google/subcommands"
)

type Config struct {
	Server ServerConfig
}

type ServerConfig struct {
	Name string
}

type TestCmd struct {
	sentence   string
	ExecTest   bool
	Configtest bool
}

func (*TestCmd) Name() string     { return "test" }
func (*TestCmd) Synopsis() string { return "test args to stdout." }
func (*TestCmd) Usage() string {
	return `test [-capitalize] <some text>:
  Print a famous words
`
}

func (p *TestCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.sentence, "T", "RX-78-2", "testtest")
	f.BoolVar(&p.ExecTest, "exec", false, "Test")
	f.BoolVar(&p.Configtest, "configtest", false, "")
}

func (p *TestCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	var config Config
	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(5 * time.Second)
	fmt.Println("GANDOM")
	if p.ExecTest {
		fmt.Printf("%+v", os.Args[0], os.Args[1])
		fmt.Println("ExecTest is Run")
		if len(os.Args) == 1 {
			// exec itself
			cmd := exec.Command("go", "run", "main.go", "test", "&")
			cmd.Start()
		} else {
			fmt.Println("aaaa")
		}

	} else if p.Configtest {
		fmt.Printf("Name is %s\n", config.Server.Name)
	}
	return subcommands.ExitSuccess
}
