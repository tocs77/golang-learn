package main

import "fmt"

func main() {

	arr := [...]int{1, 2, 3, 4, 5}

	fmt.Println("Before changing")

	slice := arr[2:]

	fmt.Println(arr)
	fmt.Println(slice, len(slice), cap(slice))

	arr[3] = 0
	slice[2] = 10

	fmt.Println("After changing")
	fmt.Println(arr)
	fmt.Println(slice, len(slice), cap(slice))

	slice = append(slice, 15)
	arr[4] = -1

	fmt.Println("After appending")
	fmt.Println(arr)
	fmt.Println(slice, len(slice), cap(slice))
}
