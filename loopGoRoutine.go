package main

import (
	"fmt"
)

func main() {
	block := make(chan bool, 1)
	block <- true
	var counter int
	for i := 0; i < 1000; i++ {
		<-block
		go func() {
			counter++
			block <- true
		}()

	}
	//time.Sleep(time.Second * 1)
	fmt.Println("counter: ", counter)
}
