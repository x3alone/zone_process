package piscine

func StringToIntSlice(str string) []int {
	var s []int
	for _, val := range str {
		s = append(s, int(val))
	}
	return s
}
