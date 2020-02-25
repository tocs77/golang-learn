package main

import "fmt"

func main() {
	nums := [...]int{10, 15, 30}

	var sum int
	nums[1] = 20

	for i := range nums {
		sum += nums[i]
	}

	fmt.Println(sum)
	fmt.Println(nums)
	fmt.Println(nums[0])
	fmt.Println(nums[2])
	// fmt.Println(nums[3]) //! Compiler error
	fmt.Println(len(nums))

}
