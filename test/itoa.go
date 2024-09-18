package main

import "fmt"

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	isneg := n<0 
	if isneg{
		n= -n
	}
	res := ""
	for n>0 {
		dig := n % 10
		res = string(dig+ '0') +res
		n /= 10
	}
	if isneg{
		res = "-" + res
	}
	return res

}

func main() {
	fmt.Println(Itoa(123))
}
