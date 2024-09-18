package main

import (
	"fmt"
)

func FromTo(from int, to int) string {
	if from > 99 || to < 0 || to > 99 || from < 0{
		return "invalid\n"
	}
	var res string
	if from < to {
		for i := from; i <= to; i++{
			if i < 10{
				res += fmt.Sprintf("0%d", i)
			}else {
				res += fmt.Sprintf("%d", i)
			}
			if i< to {
				res += ", "
			}
		}
	} else {
		for i := from; i >= to; i--{
			if i < 10{
				res += fmt.Sprintf("0%d", i)
			}else {
				res += fmt.Sprintf("%d", i)
			}
			if i> to {
				res += ", "
			}
		}
	}
	return res + "\n"
}

func main() {
	fmt.Print(FromTo(1, 10))
	fmt.Print(FromTo(10, 1))
	fmt.Print(FromTo(10, 10))
	fmt.Print(FromTo(100, 10))
}
