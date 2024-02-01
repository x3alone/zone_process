package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	a := 1
	if n < 0 {
		a = -1
		z01.PrintRune('-')
	}
	if n != 0 {
		z := (n / 10) * a
		if z != 0 {
			PrintNbr(z)
		}
		k := (n % 10 * a) + '0'
		z01.PrintRune(rune(k))
	} else {
		z01.PrintRune('0')
	}
}
