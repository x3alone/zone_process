package main

import "fmt"

func FishAndChips(n int) string {
	if n < 0 {
		return "error: number is negative"
	}

	// Check if divisible by 2 and 3
	if n%2 == 0 && n%3 == 0 {
		return "fish and chips"
	}

	// Check if divisible by 2
	if n%2 == 0 {
		return "fish"
	}

	// Check if divisible by 3
	if n%3 == 0 {
		return "chips"
	}

	// If neither divisible by 2 or 3
	return "error: non divisible"
}

func main() {
	fmt.Println(FishAndChips(0))
	fmt.Println(FishAndChips(9))
	fmt.Println(FishAndChips(6))
}
