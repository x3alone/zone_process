package mazzin

import "fmt"

func RetainFirstHalf(str string) string {
	if len(str) == 1 {
		return str
	} else if len(str)%2 == 0{
		return str[:len(str)/2]
	} else if len(str)%2 != 0{
		return str[:len(str)/2]
	} else{
		return ""
	}
}

func main() {
	fmt.Println(RetainFirstHalf("This is the 1st halfThis is the 2nd half"))
	fmt.Println(RetainFirstHalf("A"))
	fmt.Println(RetainFirstHalf(""))
	fmt.Println(RetainFirstHalf("Hello World"))
}
