package piscine

func IsLower(s string) bool {
	for _, t := range s {
		if !(t >= 97 && t <= 122) {
			return false
		}
	}
	return true
}
