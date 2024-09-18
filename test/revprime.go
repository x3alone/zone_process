package main

import "fmt"

func isprime(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func FindPrevPrime(nb int) int {
	if nb <= 2 {
		return -1
	}
	n := nb
	for n > 2 && !isprime(n) {
		n--
	}
	return n
}

func main() {
	fmt.Println(FindPrevPrime(5))
	fmt.Println(FindPrevPrime(4))
}
