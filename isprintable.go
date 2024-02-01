package piscine

func IsPrintable(s string) bool {
	for _, t := range s {
		if !(t >= 32 && t <= 126) {
			return false
		}
	}
	return true
}
