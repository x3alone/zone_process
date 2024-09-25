package main

import "fmt"

func WeAreUnique(str1, str2 string) int {
	// str3 := str1 + str2
	// var count int
	// if str3 == "" {
	// 	return -1
	// }
	// mmap := make(map[rune]int)
	// for _, l := range str3 {
	// 	mmap[l]++
	// }
	// for _, mpl := range mmap {
	// 	if mpl == 1 {
	// 		count++
	// 	} else if mpl < 2 {
	// 		return 0
	// 	}
	// }
	if str1 + str2 == ""{
		return -1
	}
	count := 0
	j := 0
	for i :=0 ; i<len(str2)&& j< len(str1) ;i++{
		if str1[j] == str2[i]{
			j++
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
