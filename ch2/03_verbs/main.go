package main

import "fmt"

func main() {
	var k1 uint8 = 20

	f1 := 3.14
	fmt.Println("Hi")
	fmt.Printf("%f - %T\n", f1, f1)
	fmt.Printf("%5.3f %.2f \n", f1, 214.437)

	fmt.Println(k1+12, k1-12, k1*2, k1/2, 5%3)

	k1++
	fmt.Println(k1)

	k1 += 10
	fmt.Println(k1)

	fmt.Println(5/2, 5.0/2.0)

}
