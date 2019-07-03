package commands

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	portscanner "github.com/anvie/port-scanner"
	"github.com/google/subcommands"
	pipeline "github.com/mattn/go-pipeline"
)

type PortCmd struct {
	Test    bool
	Scan    bool
	Monitor bool
}

func (*PortCmd) Name() string     { return "port" }
func (*PortCmd) Synopsis() string { return "Port Handling" }
func (*PortCmd) Usage() string {
	return "port [-capitalize] <some text>:"
}

func (p *PortCmd) SetFlags(f *flag.FlagSet) {
	f.BoolVar(&p.Test, "test", false, "TEST option")
	f.BoolVar(&p.Scan, "scan", false, "Usuge")
	f.BoolVar(&p.Monitor, "monitor", false, "Usage : port -monitor [MonitorPortNumber]")
}
func (p *PortCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {

	//config.TomlLoader()
	args := f.Args()

	if p.Test {
		fmt.Println("MS-06S")
	} else if p.Scan {
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
			WallNotification(TargetPort,TargetPortProcess)
			os.Exit(1)
		}
		time.Sleep(1 * time.Second)
	}
}

func WallNotification(TargetPort int,TargetPortProcess string) {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatal(err)
	}

	nowtime := fmt.Sprintf("%s", time.Now().In(loc))
	mes := fmt.Sprintf("Warning : [%d] : %s was killed at %s \n", TargetPort, TargetPortProcess, nowtime)
	_, err = pipeline.Output(
		[]string{"echo", mes},
		[]string{"wall"},
	)
	if err != nil {
		log.Fatal(err)
	}
}
