package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for i := '9'; i >= '0'; i-- {
		for l := '9'; l >= '0'; l-- {
			for k := '9'; k >= '0'; k-- {
				for j := '9'; j >= '0'; j-- {
					if i > k || (i == k && l > j) {
						z01.PrintRune(i)
						z01.PrintRune(l)
						z01.PrintRune(' ')
						z01.PrintRune(k)
						z01.PrintRune(j)
						if !(i == '0' && l == '1' && k == '0' && j == '0') {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
}
