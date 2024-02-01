package main

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for _, z := range s{
		z01.PrintRune(z)
	}
}