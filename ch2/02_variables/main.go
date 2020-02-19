package main

import "fmt"

func main() {
	var i int
	fmt.Println("i = ", i)

	j := 2
	fmt.Printf("%d - %T\n", j, j)

	var k1 uint8 = 20
	// k1 = 256 //complier error
	//k1 = -2 //compiler error

	fmt.Printf("%d - %T\n", k1, k1)

	var k2 uint16 = 30

	// k1 = k2 //compiler error
	//k2 = k1 //compiler error

	fmt.Printf("%d - %T\n", k2, k2)
}
