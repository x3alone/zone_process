package main


import "fmt"


func ThirdTimeIsACharm(str string) string {
	if str == "" || len(str) <3{
		return "\n"
	}
	res := ""
	for i := 2 ; i < len(str); i += 3{
		res += string(str[i])
	}
	return res + "\n"
}

func main() {
	fmt.Print(ThirdTimeIsACharm("123456789"))
	fmt.Print(ThirdTimeIsACharm(""))
	fmt.Print(ThirdTimeIsACharm("a b c d e f"))
	fmt.Print(ThirdTimeIsACharm("12"))
}


/*package piscine

func ThirdTimeIsACharm(str string) string {
	if str == "" || len(str) <3{
		return "\n"
	}
	strr:= len(str)
	st := ""
	for i:= 0 ; i < len(str); i++{
		for j := i+3; j < len(str); j++{
			st = string( str [i]+ str [j])
			if len(st) == strr{
				return st
			}
		}
	}
	return "\n"
}
*/



