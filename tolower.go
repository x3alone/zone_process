package piscine

func ToLower(s string) string {
	a := ""
	for _, t := range s {
		if t >= 65 && t <= 90 {
			a += string(t + 32)
		} else {
			a += string(t)
		}
	}
	return a
}
