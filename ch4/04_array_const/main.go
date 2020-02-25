package main

import "fmt"

func main() {
	type Size int

	const (
		EX Size = iota
		LG
		MD
		SM
	)

	sz := [...]string{EX: "Extra Large", LG: "Large", MD: "Medium", SM: "Small"}

	fmt.Println(MD, sz[MD])
	fmt.Println(sz)

	x := [...]int{3: 10, 1: 20}

	fmt.Println(x)
}
