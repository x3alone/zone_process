package piscine

func Index(s string, toFind string) int {
	if s == "" {
		return -1
	} else if toFind == "" {
		return 0
	}
	str := []rune(s)
	toF := []rune(toFind)

	for i := 0; i <= len(str)-1; i++ {
		match := true

		for j := 0; j < len(toF); j++ {
			if str[i+j] != toF[j] {
				match = false
				break
			}
		}

		if match {
			return i
		}
	}

	return -1
}
