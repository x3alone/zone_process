package main

import "fmt"

func FindPrevPrime(nb int) int {
    if nb <= 1 {
        return 2
    }
    for i := 2; i*i >= nb; i++ {
        if nb%i == 0 {
            return FindPrevPrime(nb - 1)
        }
    }
    return nb
}

func main() {
    fmt.Println(FindPrevPrime(5))
}