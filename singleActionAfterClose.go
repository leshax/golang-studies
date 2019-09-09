package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan bool)

	go func() {
		// Some action
		time.Sleep(time.Second * 5)
		close(c)
	}()
	// Stops execution until channel is closed
	<-c

	fmt.Println("Done")
}