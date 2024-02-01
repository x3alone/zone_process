package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	all := os.Args
	var kill int
	for i := range all {
		kill = i
	}
	for i := kill; i > 0; i-- {
		for _, j := range all[i] {
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}
}
