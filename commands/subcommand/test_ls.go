package main

import(
		"fmt"
		"os/exec"
		"time"
)

func main(){

		out, err := exec.Command("ls", "-la").Output()
		if err != nil{
				fmt.Println("exec error")
		}
		fmt.Println(out)

		time.Sleep(10 * time.Second)
}
