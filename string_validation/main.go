package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Scan(&input)

	if validateString(input) {
		fmt.Println("TRUE")
	} else {
		fmt.Println("FALSE")
	}
}

func validateString(input string) bool {
	stack := []rune{}
	matchingBrackets := map[rune]rune{')': '(', '}': '{', ']': '[', '>': '<'}

	for _, char := range input {
		if char == '(' || char == '{' || char == '[' || char == '<' {
			stack = append(stack, char)
		} else if closing, ok := matchingBrackets[char]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != closing {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
