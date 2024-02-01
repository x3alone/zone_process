package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	var kok string
	var sl []string
	for _, char := range str {
		if char != ' ' {
			kok += string(char)
		} else {
			sl = append(sl, kok)
			kok = ""
		}
	}
	sl = append(sl, kok)
	summary := make(map[string]int)
	for _, kok := range sl {
		summary[kok]++
	}
	return summary
}
