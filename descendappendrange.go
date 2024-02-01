package piscine

func DescendAppendRange(max, min int) []int {
	if max <= min {
		return []int{}
	}
	var result []int
	for i := max; i > min; i-- {
		if min > max {
			return result
		}
		result = append(result, i)
	}
	return result
}
