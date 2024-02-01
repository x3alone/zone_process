package piscine

func AlphaCount(s string) int {
	t := []rune(s)
	count := 0
	for i := 0; i <= len(s)-1; i++ {
		if (t[i] >= 'a' && t[i] <= 'z') || (t[i] >= 'A' && t[i] <= 'Z') {
			count++
		}
	}
	return count
}
