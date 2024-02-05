package main

import "fmt"

func DigitLen(n, base int) int {
	if !(base >= 2 && base <= 36){
	return -1
	}
	if n == 0{
		return 0
	}
	counter := 0
	neg := -1
	if n < 0{
		n *= neg
	}
	for n > 0 {
		n/= base
		counter ++
		if n / base == 0{
			counter++ 
			break
		}
	}
	return counters
}

func main() {
	fmt.Println(DigitLen(100, 10))
	fmt.Println(DigitLen(100, 2))
	fmt.Println(DigitLen(-100, 16))
	fmt.Println(DigitLen(100, -1))
}