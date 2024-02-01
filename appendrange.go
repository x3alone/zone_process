package piscine

func AppendRange(min, max int) []int {
	var result []int
	for i := 0; i < max; i++ {
		if min > max {
			return result
		}
		result = append(result, i)
	}
	return result
}
