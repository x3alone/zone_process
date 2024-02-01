package piscine

func CountIf(f func(string) bool, tab []string) int {
	count := 0
	for _, i := range tab {
		if f(i) {
			count += 1
		}
	}
	return count
}
