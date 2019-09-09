package main

import (
	"fmt"
	"time"
)

func background(die chan bool, value int) {
	for {
		select {
		default:
			fmt.Println("value: ", value)
			time.Sleep(time.Second * 1)
		case <-die:
			fmt.Println("Go routine closed: ", value)
			die <- true
			return
		}
	}
}

func main() {
	die := make(chan bool)
	go background(die, 1)
	go background(die, 2)
	time.Sleep(time.Second * 5)
	die <- true
	//Wait...
	time.Sleep(time.Second * 5)
	fmt.Println("End.")
}
