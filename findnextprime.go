package piscine

func FindNextPrime(nb int) int {
	if nb == 3 {
		return nb
	} else if nb > 3 {
		for {
			prime := true
			for i := 2; i*i <= nb; i++ {
				if nb%i == 0 {
					prime = false
					break
				}
			}
			if prime {
				return nb
			}
			nb++
		}
	}
	return 2
}
