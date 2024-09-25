package main

import (
	"os"
	"github.com/01-edu/z01"
)
// func split(s string) []string {
// 	var strs []string 
// 	word:= ""
// 	for ind, val := range s{
// 		if val != ' ' || val != '\r' || val != '\n'{
// 			word += string(val)
// 		}else {
// 			strs = append(strs, word)
// 			word = ""
// 		}
// 	}
// 	return strs
// }

// func string(s string){
// 	for _, c := range s{
// 		z01.PrintRune(c)
// 	}
// }

// func main(){
// 	if len(os.Args) != 2 {
// 		return
// 	}
// 	str := os.Args[1]
// 	strs := split(str)
	
// }





func split(s string)[]string{
	var all []string
	word := ""
	for i := 0 ; i < len(s); i++{
		if s[i] != ' ' && s[i] != '\n' && s[i] != '\r'{
			word += string(s[i])
		} else if word != ""{
			all = append(all, word)
			word = ""
		}
	}
	if word != ""{
		all = append(all, word)
	}
	return all
}
func print(s string){
	for _,i :=range s{
		z01.PrintRune(i)
	}
}

func main(){
	if len(os.Args) != 2{
		return
	}
	str := split(os.Args[1])
	b :=0 
	for _,i := range str{
		print(i)
		if b < len(str)-1 {
			print("   ")
		}
	}
	print("\n")
}









// func split(s string) []string {
// 	var so []string
// 	ss:= ""
// 	for ind , val := range s{
// 		if val != ' ' || val != '\n' || val != '\r'{
// 			ss += string(val)
// 		} else (val == ' ' || val == '\r' || val == '\n') && ind == len(s)-1 || ss != '' {
// 			so = append(so, ss)
// 			ss = ""
// 		}
// 	}
// 	return so
// }

// func main(){

// 	if len(os.Args) != 2 {
// 		return
// 	}
// 	str := os.Args[1]

// 	for ind, val := range str{
		 
// 	}
// }