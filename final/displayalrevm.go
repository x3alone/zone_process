package main

import (
	"fmt"
)

func main() {
	for i := 25; i >= 0; i-- {
		char := rune('A' + i)
		if i%2 == 1 {
			char = rune('a' + i)
		}
		fmt.Printf("%c", char)
	}
	fmt.Println()
}
