package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(myPow(2, -2))
}
func myPow(x float64, n int) float64 {
	return Pow(x, n)
}

func Pow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}
	if n%2 == 0 {
		half := myPow(x, n/2)
		return half * half
	} else {
		return x * myPow(x, n-1)
	}
}
