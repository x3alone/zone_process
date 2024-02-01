package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	last := os.Args[0]
	for i, p := range last {
		if i > 1 {
			z01.PrintRune(p)
		}
	}
	z01.PrintRune('\n')
}
