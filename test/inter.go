package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return 
	}

	str1 := os.Args[1]
	str2 := os.Args[2]

	seen := make(map[rune]bool)
	var result []rune

	for _, char := range str1 {
		if !seen[char] && contains(str2, char) {
			result = append(result, char)
			seen[char] = true
		}
	}

	fmt.Println(string(result))
}

func contains(str string, char rune) bool {
	for _, c := range str {
		if c == char {
			return true
		}
	}
	return false
}