package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestValidParentheses("()"))
}
func longestValidParentheses(s string) int {
	stack := []int{-1}
	maxLength := 0

	for i, char := range s {
		if char == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				currentLength := i - stack[len(stack)-1]
				if currentLength > maxLength {
					maxLength = currentLength
				}
			}
		}
	}

	return maxLength
}
