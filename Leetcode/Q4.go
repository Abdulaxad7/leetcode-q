package main

import "fmt"

func main() {
	fmt.Println(isValid("()[]{}"))
}

func isValid(s string) bool {
	stack := []rune{}
	m := map[rune]rune{')': '(', '}': '{', ']': '['}
	for _, c := range s {
		switch c {
		case '(', '{', '[':
			stack = append(stack, c)
		case ')', '}', ']':
			if len(stack) == 0 || stack[len(stack)-1] != m[c] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
