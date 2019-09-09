package main

import "fmt"

func main() {
	id := make(chan string)

	go func() {
		var counter int64 = 1
		for {
			//fmt.Println("Generating...")
			id <- fmt.Sprintf("%x", counter)
			counter += 1
		}
	}()

	fmt.Printf("%s\n", <-id) // will be 1
	fmt.Printf("%s\n", <-id) // will be 2
	fmt.Printf("%s\n", <-id) // will be 3
	fmt.Printf("%s\n", <-id) // will be 4
	fmt.Printf("%s\n", <-id) // will be 5
}
