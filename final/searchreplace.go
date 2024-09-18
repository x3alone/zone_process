package main

import (
	"os"

	"github.com/01-edu/z01"
)


func main() {
	if len(os.Args) != 4 {
		return
	}
	str := os.Args[1]
	tr := os.Args[2]
	re := os.Args[3]
	if len(tr) != 1 {
		return
	}
	res := []rune{}
	for _, char := range str {
		if string(char) == tr {
			res = append(res, rune(re[0]))
		} else {
			res = append(res , rune(char))
		}
	}
	for _, i := range res{
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}