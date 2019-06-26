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

		WallOut, err := exec.Command("wall","TEST").Output()
		if err != nil{
				fmt.Println(err)
				fmt.Println("occerd error")
				//break
		}
		fmt.Println(WallOut)
}
