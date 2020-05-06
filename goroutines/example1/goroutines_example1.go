package main

import (
	"fmt"
	"time"
)

func main() {

	go printHello()
	//Wait for the go routine to complete
	time.Sleep(2 * time.Second)
	//Wait over, print and exit
	fmt.Println("Completed Waiting Exiting...")
}

func printHello() {

	fmt.Println("Hello from the printHello function !")
}
