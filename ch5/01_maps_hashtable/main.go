package main

import "fmt"

func main() {
	values := []string{"ABC", "ACB", "BCA", "CAB", "CBA"}

	factor := []int{100, 10, 1}

	//* 65*100 + 66*10 + 67*1 = 7227

	hashKey := 0

	for index := range values {
		fmt.Printf("Incoming value: %s  ", values[index])
		bytes := []byte(values[index])
		f := 0
		for i := range bytes {
			fmt.Print(bytes[i], " ")
			hashKey += int(bytes[i]) * factor[f]
			f++
		}
		fmt.Printf("(hashKey: %d) \n", hashKey)
	}
}
