package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	scanner := bufio.NewScanner(os.Stdin)
	stringsList := make([]string, n)

	// Reading strings
	for i := 0; i < n; i++ {
		scanner.Scan()
		stringsList[i] = scanner.Text()
	}

	result := findMatchingStrings(stringsList)

	if len(result) == 0 {
		fmt.Println("false")
	} else {
		// Print the first set of matched indices
		for i, idx := range result {
			if i == len(result)-1 {
				fmt.Printf("%d", idx+1)
			} else {
				fmt.Printf("%d ", idx+1)
			}
		}
		fmt.Println()
	}
}

func findMatchingStrings(stringsList []string) []int {
	for i := 0; i < len(stringsList); i++ {
		for j := i + 1; j < len(stringsList); j++ {
			// Check for case-insensitive equality
			if strings.EqualFold(stringsList[i], stringsList[j]) {
				return []int{i, j}
			}
		}
	}
	return nil
}
