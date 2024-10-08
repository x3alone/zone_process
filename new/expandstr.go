package main

import (
	"os"
	"github.com/01-edu/z01"
)

func split(s string)[]string{
	var res []string
	word:=""
	for i:= 0 ; i < len(s); i++{
		if s[i] != ' ' && s[i] != '\r' && s[i] != '\n'{
			word+=string(s[i])
		} else if word != ""{
			res=append(res, word)
			word =""
		}
	}
	if word != ""{
		res=append(res, word)
	}
	return res
}

func print(s string) {
	for _,o := range s{
		z01.PrintRune(o)
	}
}

func main() {
	if len(os.Args) != 2{
		return
	}
	
	words := split(os.Args[1])
	i := 0
	for _, in:= range words{
		print(in)
		if i < len(words)-1{
			print("   ")
		}
	}
	print("\n")
}

// func split(str string) []string {
// 	var res []string
// 	word := ""
// 	for _, val := range str {
// 		if val != ' ' && val != '\n' && val != '\r' {
// 			word += string(val)
// 		} else if word != "" {
// 			res = append(res, word)
// 			word = ""
// 		}
// 	}
// 	if word != "" {
// 		res = append(res, word)
// 	}
// 	return res
// }
// func printstring(s string) {
// 	for _,i := range s {
// 		z01.PrintRune(i)
// 	}
// }


// func main() {
// 	if len(os.Args) != 2 {
// 		return
// 	}
// 	str := os.Args[1]
// 	words := split(str)

// 	if len(words) == 0 {
// 		return
// 	}
// 	i := 0
// 	for _, word := range words {
// 		printstring(word)
// 		if i < len(words)-1 {
// 			printstring("   ")
// 		}
// 	}
// 	printstring("\n")
// }
