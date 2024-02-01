package piscine

func NRune(s string, n int) rune {
	if s == "" || n <= 0 || n > len(s) {
		return 0
	}
	t := []rune(s)
	return t[n-1]
}
