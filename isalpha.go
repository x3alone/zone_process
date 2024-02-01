package piscine

func IsAlpha(s string) bool {
	for _, t := range s {
		if !((t >= 97 && t <= 122) || (t >= 65 && t <= 90) || (t >= 48 && t <= 57)) {
			return false
		}
	}
	return true
}
