package main

import (
	"fmt"
	"time"
)

func pause(c chan bool) {
	// --- Pause ---
	<-c
	// --- Pause ---
	fmt.Println("GO_RUN")
}

func main() {
	c := make(chan bool)
	fmt.Println("Run go routines...")
	for i := 0; i < 3; i++ {
		go pause(c)
	}
	fmt.Println("Pausing...")
	time.Sleep(time.Second * 5)
	fmt.Println("Close chanel...")
	close(c)
	time.Sleep(time.Second * 5)
}
