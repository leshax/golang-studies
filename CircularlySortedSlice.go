package main

import "fmt"

func IsCircleSorted(r []int) bool {

	fmt.Println("!")

	if len(r) == 1 {
		return true
	}

	splitIndex := -1

	for i := 0; i < len(r)-1; i++ {
		if r[i] < r[i+1] {
			splitIndex := i
			return
		}
	}

	if splitIndex == -1 {
		return true
	}

	fmt.Println(" ", count)
	if count >= (len(r) - 2) {

		return true
	} else {
		return false
	}

}

func main() {
	v1 := IsCircleSorted([]int{2, 3, 4, 5, 6})
	v2 := IsCircleSorted([]int{6, 2, 3, 4, 5})
	v3 := IsCircleSorted([]int{3, 2, 4, 5, 6})
	fmt.Println("v1: ", v1)
	fmt.Println("v2: ", v2)
	fmt.Println("v3: ", v3)
}
