package main

import (
	"os"

	"github.com/01-edu/z01"
)
func print(s string) string{
	for _ ,i := range s{
		z01.PrintRune(i)
	}
	return s
}
func main(){
	if len(os.Args) != 2 || os.Args[1] == ""{
		z01.PrintRune('\n')
		return
	}
	input := os.Args[1]

	var res []byte
	innit := false
	for i := 0 ; i< len(input); i++{
		char := input[i]
		if char == ' ' || char == '\t'{
			if innit{
				innit = false
				res = append(res, ' ')
			}
		}else {
			res = append(res, char)
			innit = true
		}
	}
	if len(res) > 0 && res[len(res)-1] == ' '{
		res = res[:len(res)-1]
		}
		print(string(res))
		z01.PrintRune('\n')

}