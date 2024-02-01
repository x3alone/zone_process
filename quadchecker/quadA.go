package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func QuadA(x, y int) {
	if x <= 0 {
		return
	}
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if i == 0 || i == y-1 {
				if j == 0 || j == x-1 {
					z01.PrintRune('o')
				} else {
					z01.PrintRune('-')
				}
			} else {
				if j == 0 || j == x-1 {
					z01.PrintRune('|')
				} else {
					z01.PrintRune(' ')
				}
			}
		}
		z01.PrintRune('\n')
	}
}

func main() {
	args := os.Args
	if len(args) == 3 {
		x, err1 := strconv.Atoi(args[1])
		y, err2 := strconv.Atoi(args[2])
		if x == 0 || y == 0 {
			fmt.Println("Not a quad function")
			return
		}
		if err1 == nil && err2 == nil {
			QuadA(x, y)
		} else {
			fmt.Println("Not a quad function")
		}
	} else {
		fmt.Println("Not a quad function")
	}
}
