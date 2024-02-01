package piscine

func IsNumeric(s string) bool {
	for _, t := range s {
		if !(t >= 48 && t <= 57) {
			return false
		}
	}
	return true
}
