package main

import "fmt"

func main() {
	var k1 uint8 = 20

	const pi = 3.14159

	// pi = 3.2 // error cant assign to const

	k1 /= 2

	fmt.Println(25, pi, k1)

	const (
		a = 2
		b
		c
		d = a * 10
		e
		f
	)

	fmt.Println(a, b, c, d, e, f)

	const (
		a2 = iota + 1
		b2
		c2
		d2 = a2 * 10
		e2
		f2 = iota
	)
	fmt.Println(a2, b2, c2, d2, e2, f2)

	const (
		a3 = iota + 1 //00000001
		b3 = a3 << 1  //00000010
		c3 = b3 << 1
		d3 = c3 << 1
		e3 = d3 << 1
	)

	fmt.Println(a3, b3, c3, d3, e3)
}
