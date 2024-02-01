package piscine

func ToUpper(s string) string {
	a := ""
	for _, t := range s {
		if t >= 97 && t <= 122 {
			a += string(t - 32)
		} else {
			a += string(t)
		}
	}
	return a
}
