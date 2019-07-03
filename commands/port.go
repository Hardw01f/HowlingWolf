package commands

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"runtime"
	"strconv"
	"sync"
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

var Dist string = "localhost"

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
			if flag.NArg() != 4 {
				fmt.Println("Usage : port -scan [Start-port] [End-port]")
				os.Exit(1)
			}

			StartPort, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println(err)
			}

			EndPort, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println(err)
			}

		wg := new(sync.WaitGroup)
		cpus := runtime.NumCPU()
		fmt.Printf("runtime.CPUS : %d\n", cpus)
		runtime.GOMAXPROCS(cpus)

		var semaphoMultiple int
		if cpus == 2 {
			semaphoMultiple = 150
		} else {
			semaphoMultiple = 300
		}
		fmt.Printf("semaphoMultiple : %d\n", semaphoMultiple)

		semapho := make(chan int, cpus*semaphoMultiple)

		for {
			logfile, err := os.OpenFile("portscan.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
			if err != nil {
				log.Fatal(err)
			}
			defer logfile.Close()

			jpn, err := time.LoadLocation("Asia/Tokyo")
			nowtime := time.Now().In(jpn)
			fmt.Fprintf(logfile, " --- %s --- \n", nowtime)

			for PortNum := StartPort; PortNum <= EndPort; PortNum++ {
				//fmt.Printf("Goroutine : %d\n", runtime.NumGoroutine())
				//fmt.Println(PortNum)
				/*if PortNum == Maxnum {
						fmt.Println("last port started")
				}*/

				wg.Add(1)
				go Scan(PortNum, semapho, wg, logfile)
			}
			wg.Wait()
			fmt.Println("finish")
			fmt.Fprintln(logfile, "")
			//DebugScan(5060)
			time.Sleep(180 * time.Second)
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
			WallNotification(TargetPort, TargetPortProcess)
			os.Exit(1)
		}
		time.Sleep(1 * time.Second)
	}
}

func WallNotification(TargetPort int, TargetPortProcess string) {
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

func Scan(PortNum int, semapho chan int, wg *sync.WaitGroup, file *os.File) {
	semapho <- 1

	_, err := net.DialTimeout("tcp", Dist+":"+strconv.Itoa(PortNum), (1500 * time.Millisecond))

	if err != nil {
		/*fmt.Printf(" Checking %d Port\n",PortNum)
		if PortNum%100 ==0{
				fmt.Printf("Goroutine num is       %d\n",runtime.NumGoroutine())
		}*/

	} else {
		fmt.Fprintf(file, "%d Port Opened\n", PortNum)
		fmt.Printf("%d  Port Opened\n", PortNum)
		//fmt.Printf("Goroutine : %d\n", runtime.NumGoroutine())

	}
	<-semapho
	defer wg.Done()
}
