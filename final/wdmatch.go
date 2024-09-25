package main

import (
	"os"
)

// func canform(first, second string) bool {
// 	j := 0
// 	for i := 0; i < len(second) && j < len(first); i++ {
// 		if first[j] == second[i] {
// 			j++
// 		}
// 	}
// 	return j == len(first)
// }



func canform(first, second string)bool {
	j := 0
	for i :=0 ; i<len(second)&& j< len(first) ;i++{
		if first[j] == second[i]{
			j++
		}
	}
	return j== len(first)
}






func main() {
	if len(os.Args) != 3 {
		return
	}
	if canform(os.Args[1], os.Args[2]) {
		println(os.Args[1])
	}
}
