package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 26; i++ {
		char := rune('A' + i)
		if i%2 == 0 {
			char = rune('a' + i)
		}
		fmt.Printf("%c", char)
	}
	fmt.Println()
}
