package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func prt(s string) {
	for _, i := range s {
		z01.PrintRune(i)
	}
}

func ZipString(s string) string{
	res:= ""
	for i := 0; i < len(s); i++{
		ch := s[i]
		count := 1
		for i+1< len(s)&& s[i+1] == ch{
			count++
			i++
		}
		res += string(count+ '0') + string(ch)
	}
	return res
}

// func ZipString(s string) string {
// 	res := ""
// 	n := len(s)
// 	for i := 0; i < n; i++ {
// 		ch := s[i]
// 		count := 1
// 			for i+1 < n && s[i+1] == ch {
// 				count++
// 				i++
// 			}
// 			res += string(count+'0') + string(ch)
		
// 	}
// 	return res
// }

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}

// hello
// 1h1e2l1o
