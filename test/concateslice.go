package main

import "fmt"

func ConcatSlice(slice1, slice2 []int) []int {
	res := make([]int, len(slice1)+len(slice2))
	copy(res, slice1)
	copy(res[len(slice1):], slice2)
	return res
}

func main() {
	fmt.Println(ConcatSlice([]int{}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{}))
}
