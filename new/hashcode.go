package main

import (
	"fmt"
)

func HashCode(dec string) string {
	var res []rune
	for _, i := range dec{
		new:=(int(i)+len(dec)) %127
		if new< 33{
			new += 33
		}
		res = append(res, rune(new))
	}
	return string(res)
}

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}

//abc
//def
