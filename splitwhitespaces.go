package piscine

func SplitWhiteSpaces(s string) []string {
	var result []string
	str := []rune(s)
	NewStr := ""
	for i := 0; i < len(s); i++ {
		if (str[i] == 32) || (str[i] >= 9 && str[i] <= 13) {
			if NewStr != "" {
				result = append(result, NewStr)
				NewStr = ""
			}
		} else {
			NewStr += string(str[i])
		}
	}
	if NewStr != "" {
		result = append(result, NewStr)
	}
	return result
}
