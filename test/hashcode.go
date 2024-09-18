package main

import "fmt"

func HashCode(dec string) string {
	var result []rune
	len := len(dec)
	for _,i := range dec {
		new := (int(i)+len) % 127
		if new < 33{
			new += 33
		}
		result = append(result, rune(new))
	}
	return string(result)
}

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}
