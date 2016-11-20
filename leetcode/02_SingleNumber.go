package main

import "fmt"

func main() {
	nums := []int{1, 1, 2, 2, 3, 3, 4}

	fmt.Println("으아.")
	fmt.Println("SingleNumber..", singleNumber(nums))
}

func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	for i := 1; i < len(nums); i++ {
		nums[0] ^= nums[i]
	}

	return nums[0]
}