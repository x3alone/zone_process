package piscine

func Atoi(s string) int {
	t := 0
	n := 1

	for ind, val := range s {
		if ind == 0 && (val == '-' || val == '+') {
			if val == '-' {
				n = -1
			}
		} else if val >= '0' && val <= '9' {
			t = t*10 + int(val-'0')
		} else {
			return 0
		}
	}
	return t * n
}