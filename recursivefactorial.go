package piscine

// import "fmt"
func RecursiveFactorial(nb int) int {
	if nb < 0 || nb >= 21 {
		return 0
	}
	if nb == 0 {
		return 1
	}
	return nb * RecursiveFactorial(nb-1)
}

/*
func main() {
	arg := 12
	fmt.Println(RecursiveFactorial(arg))
}*/
