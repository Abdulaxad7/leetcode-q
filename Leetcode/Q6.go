package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1, 2, 4, 5, 6}, 5))
}

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		midPoint := (left + right) / 2
		if nums[midPoint] < target {
			left = midPoint + 1
		} else if nums[midPoint] > target {
			right = midPoint - 1
		} else {
			return midPoint
		}
	}
	return left
}
