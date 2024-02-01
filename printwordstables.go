package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, lee := range a {
		for _, a := range lee {
			z01.PrintRune(a)
		}
		z01.PrintRune('\n')
	}
}
