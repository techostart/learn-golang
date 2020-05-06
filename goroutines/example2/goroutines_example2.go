package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i < 5; i++ {
		go printHello(i)
	}
	//Wait for the go routines to complete
	time.Sleep(2 * time.Second)
	//Wait over, print and exit
	fmt.Println("Completed Waiting Exiting...")
}

func printHello(id int) {

	fmt.Printf("Routine %d : Hello from the printHello function !\n", id)
}
