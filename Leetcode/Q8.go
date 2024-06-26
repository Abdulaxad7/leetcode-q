package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 5, 9, 1, 5, 9}, 2, 3))
}
func getBucketID(n, bs int) int {
	if n >= 0 {
		return n / bs
	}
	return (n+1)/bs - 1
}

func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
	if indexDiff <= 0 || valueDiff < 0 {
		return false
	}

	bs := valueDiff + 1
	b := make(map[int]int)

	for i, n := range nums {
		bi := getBucketID(n, bs)
		if _, f := b[bi]; f {
			return true
		}
		if ni, found := b[bi-1]; found && math.Abs(float64(n-ni)) <= float64(valueDiff) {
			return true
		}
		b[bi] = n
		if i >= indexDiff {
			delete(b, getBucketID(nums[i-indexDiff], bs))
		}
	}

	return false
}
