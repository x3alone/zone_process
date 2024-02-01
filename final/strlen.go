package main

import (
	"fmt"
)

func Strlen(s string) int {
	x := 0
	a:= []rune (s)
	for i := range a{
		x = i+1
	}
	return x
}

func main() {
	fmt.Println(Strlen("hahah"))
}
