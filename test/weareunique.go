package main

import "fmt"

func WeAreUnique(str1, str2 string) int {
	strr := str1 + str2
	if len(str1) == 0 && len(str2) == 0 {
		return -1
	}
	if str1 == str2 {
		return 0
	}
	count := 0
	lmap := make(map[rune]int)
	for _,i := range strr{
		lmap[i]++
	}
	for _,l:= range lmap{
		if l ==1 {
			count ++
		}
	}
	return count
}

func main() {
	fmt.Println(WeAreUnique("foo", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
}
