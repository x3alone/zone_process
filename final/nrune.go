package piscine

func NRune(s string, n int) rune {
	str :=[] rune (s)
	ln := 0
	for id := range str{
		ln = id 
	}
	if n > 0 && n < ln{
		return str[n]
	}
	return 0
}
