package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func PutStr(str string) {
	for _, c := range str {
		z01.PrintRune(c)
	}
}

func PrintNumb(n int) {
	i := '0'
	for j := 1; j <= n%10; j++ {
		i++
	}
	for j := -1; j >= n%10; j-- {
		i++
	}
	if n/10 != 0 {
		PrintNumb(n / 10)
	}
	z01.PrintRune(i)
	return
}

func main() {
	points := &point{}

	setPoint(points)

	PutStr("x = ")
	PrintNumb(points.x)
	PutStr(", y = ")
	PrintNumb(points.y)
	z01.PrintRune('\n')
}
