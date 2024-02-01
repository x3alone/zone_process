package piscine

func ActiveBits(n int) int {
	count := 0

	for n != 0 {
		if n%2 == 1 {
			count += 1
		}
		n /= 2
	}

	return count
}
