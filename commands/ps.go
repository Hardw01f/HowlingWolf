package commands

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/mitchellh/go-ps"

	"github.com/google/subcommands"
)

type PsCmd struct {
	Sentence   string
	Option     bool
	PsShow     bool
	PsMonitor  bool
	Background bool
}

func (*PsCmd) Name() string     { return "ps" }
func (*PsCmd) Synopsis() string { return "Handling Process" }
func (*PsCmd) Usage() string {
	return "ps [-capitalize] <some text>:"
}

func (p *PsCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.Sentence, "t", "RX-78-2", "testtest")
	f.BoolVar(&p.Option, "o", false, "X105 Strike")
	f.BoolVar(&p.PsShow, "show", false, "show Process list like linux ps command")
	f.BoolVar(&p.PsMonitor, "monitor", false, "Monitor specific processes by PID")
	f.BoolVar(&p.Background, "b", false, "Running Background")
}

func (p *PsCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	fmt.Println("GANDOM")

	args := f.Args()

	if p.Option {
		fmt.Println("Strike Freedom")
	} else if p.PsShow {
		fmt.Println("Starting Process")

		//pid := os.Getpid()
		processes, err := ps.Processes()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("PID    PPID       NAME")

		for _, process := range processes {
			resstr := fmt.Sprintf("%v", process)
			trimstr := strings.Split(resstr, "{")
			trimedstr := strings.Split(trimstr[1], "}")
			res := trimedstr[0]
			//fmt.Println(res)

			SplitRes := strings.Split(res, " ")
			if runtime.GOOS == "darwin" {
				fmt.Printf("%s  %s          %s \n", SplitRes[0], SplitRes[1], SplitRes[2])
			} else if runtime.GOOS == "linux" {
				fmt.Printf("%s  %s          %s \n", SplitRes[0], SplitRes[1], SplitRes[5])
			} else {
				fmt.Println("Do not support your OS")
			}
		}

	} else if p.PsMonitor {
		fmt.Println("monitor starting")

		fmt.Println(flag.NArg())
		if flag.NArg() == 2 {
			fmt.Println("Please show help : -h")
			os.Exit(1)
		}

		TargetPid, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}

		SearchRes, err := ps.FindProcess(TargetPid)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(SearchRes)

	} else if p.Background {
		fmt.Println("Running Background")
		err := exec.Command("go", "run", "./commands/subcommand/test_ls.go").Start()
		if err != nil {
			fmt.Println("out exec error")
			fmt.Println(err)
		}
	} else {
		fmt.Println("Exit")
	}
	return subcommands.ExitSuccess
}
