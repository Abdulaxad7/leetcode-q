package main

import (
	"fmt"
)

func main() {
	reverseWords("asfdsa")

}
func reverseWords(s string) string {
	var stri [][]string

	for _, str := range s {
		if str == ' ' {

		}
		stri[] = append(stri[], string(str))
	}
	fmt.Print(stri)
	return s
}
