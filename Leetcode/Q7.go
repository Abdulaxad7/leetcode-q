package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 2, 3, 4, 7}, []int{5, 6, 7, 8}))
}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var joint []int
	var result float64

	joint = nums1
	for _, v := range nums2 {
		joint = append(joint, v)

	}
	for _, v := range joint {
		result += float64(v)
	}

	sort.Ints(joint)

	if len(joint)%2 == 0 {
		result = float64(joint[len(joint)/2]+joint[len(joint)/2-1]) / 2

	} else {
		result = float64(joint[len(joint)/2])
	}

	return result
}
