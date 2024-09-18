package main

import "fmt"

func ZipString(s string) string {
	charcount := make(map[rune]int)
	var res []rune

	for _, char := range s {
		if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' {
			charcount[char]++
		}
	}
	for _, char := range s {
		if count, found := charcount[char]; found {
			res = append(res, rune(count+'0'))
			res = append(res, char)
			delete(charcount, char)
		}
	}
	return string(res)
}

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}
