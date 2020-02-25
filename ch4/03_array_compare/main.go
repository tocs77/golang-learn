package main

import "fmt"

func main() {
	//x1 := []int{10, 15, 30}

	x2 := [...]int{10, 15, 30}

	x3 := [3]int{10, 15, 30}

	var x4 [3]int = [3]int{10, 15, 30}

	fmt.Println(x2 == x3, x2 == x4, x3 == x4)
}
