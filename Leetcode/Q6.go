package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(myPow(2, -2))
}
func myPow(x float64, n int) float64 {
	return math.Pow(x, float64(n))
}
