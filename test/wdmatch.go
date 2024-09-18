package main

import (
	"fmt"
	"os"
)

func main(){

	if len(os.Args) != 3{
		return
	}
	s1 := os.Args[1]
	s2 := os.Args[2]
	if canForm(s1,s2){
		fmt.Println(s1)
	}
}

func canForm(s1,s2 string) bool{
	i := 0 
	for _, char := range s2{
		if i < len(s1) && char == rune(s1[i]){
			i++
		} 
		if i == len(s1){
			return true
		}
	}
	return false
}








