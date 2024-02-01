package piscine

func Fibonacci(index int) int {
	if index < 0 {
		return -1
	} else if index == 0 {
		return 1
	}
	a, b := 0, 1
	for i := 2; i <= index; i++ {
		a, b = b, a+b
	}
	return b
}
