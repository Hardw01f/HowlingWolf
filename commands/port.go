package commands

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"

	portscanner "github.com/anvie/port-scanner"
	"github.com/google/subcommands"
)

type PortCmd struct {
	Scan       bool
	Monitor    bool
	BackGround bool
}

func (*PortCmd) Name() string     { return "port" }
func (*PortCmd) Synopsis() string { return "Port Handling" }
func (*PortCmd) Usage() string {
	return "port [-capitalize] <some text>:"
}

func (p *PortCmd) SetFlags(f *flag.FlagSet) {
	f.BoolVar(&p.Scan, "scan", false, "Usuge")
	f.BoolVar(&p.Monitor, "monitor", false, "")
	f.BoolVar(&p.BackGround, "b", false, "run in background")
}
func (p *PortCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {

	//config.TomlLoader()
	args := f.Args()

	fmt.Println("MS-06S")
	if p.Scan {
		//fmt.Println(flag.NArg())
		if flag.NArg() != 4 {
			fmt.Println("Usage : port -scan [Start-port] [End-port]")
			os.Exit(1)
		}
		fmt.Println("Scan Starting")
		//fmt.Println(args)

		ps := portscanner.NewPortScanner("localhost", 2*time.Second, 5)

		fmt.Printf("scanning port %s-%s...\n", args[0], args[1])

		StartPort, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(err)
		}

		EndPort, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println(err)
		}

		openedPorts := ps.GetOpenedPort(StartPort, EndPort)

		for i := 0; i < len(openedPorts); i++ {
			port := openedPorts[i]
			fmt.Print(" ", port, " [open]")
			fmt.Println("  -->  ", ps.DescribePort(port))
		}
	} else if p.BackGround && p.Monitor {
		fmt.Println("backgrund")
		/*
			err := exec.Command("./deamon/deamontest").Start()
			if err != nil {
				fmt.Println(err)
			}
			//fmt.Printf("name is %s\n", config.Server.Name)

		*/
	} else if p.Monitor {
		if flag.NArg() != 3 {
			fmt.Println("Usage : port -monitor [MonitorPortNumber]")
			os.Exit(1)
		}
		CheckPort(args[0])
	}
	return subcommands.ExitSuccess
}

func CheckPort(Port string) {
	ps := portscanner.NewPortScanner("localhost", 2*time.Second, 5)

	OnePort, err := strconv.Atoi(Port)
	if err != nil {
		fmt.Println(err)
	}

	openport := ps.GetOpenedPort(OnePort, OnePort)
	if len(openport) == 0 {
		fmt.Printf("[Warnig] : %d portprocess is not found\n", OnePort)
		os.Exit(1)
	}
	//fmt.Println(openport[0], ps.DescribePort(openport[0]))
	TargetPort := openport[0]
	TargetPortProcess := ps.DescribePort(openport[0])
	fmt.Println(TargetPort, TargetPortProcess)

	for {
		openport = ps.GetOpenedPort(OnePort, OnePort)
		if len(openport) == 0 {
			str := fmt.Sprintf("Warning : [%d] : %s was killed!!\n", TargetPort, TargetPortProcess)
			err := exec.Command("wall", "Warning : ", str)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("Warnig!!! : [%d] : %s  was killed!!\n", TargetPort, TargetPortProcess)
			os.Exit(1)
		}
		//fmt.Println(openport[0], ps.DescribePort(openport[0]))
		time.Sleep(1 * time.Second)
	}
}
