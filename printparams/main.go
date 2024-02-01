package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	all := os.Args
	for i := 1; i < len(all); i++ {
		for _, j := range all[i] {
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}
}
