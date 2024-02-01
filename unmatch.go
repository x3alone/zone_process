package piscine

func Unmatch(arr []int) int {
	for _, res := range arr {
		seen := 0
		for _, rn := range arr {
			if rn == res {
				seen++
			}
		}
		if seen == 1 || seen%2 == 1 {
			return res
		}
	}
	return -1
}
