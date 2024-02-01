package piscine

func IsUpper(s string) bool {
	for _, t := range s {
		if !(t >= 65 && t <= 90) {
			return false
		}
	}
	return true
}
