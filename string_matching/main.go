package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Scan(&n)

	scanner := bufio.NewScanner(os.Stdin)
	stringsList := make([]string, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		stringsList[i] = scanner.Text()
	}

	result := findMatchingStrings(stringsList)

	if len(result) == 0 {
		fmt.Println("false")
	} else {
		// Print the result with space between numbers without trailing space
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
	indexMap := make(map[string][]int)

	// Collect all matching strings (case-sensitive)
	for i, str := range stringsList {
		indexMap[str] = append(indexMap[str], i)

		// If this string has occurred more than once, return all its indices
		if len(indexMap[str]) > 1 {
			return indexMap[str]
		}
	}

	return nil
}
