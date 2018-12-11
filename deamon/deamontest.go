package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("this is deamon")
	for {
		time.Sleep(5 * time.Second)
	}
}
