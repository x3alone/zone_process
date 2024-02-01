package piscine

func Join(strs []string, sep string) string {
	res := strs[0]
	for _, l := range strs[1:] {
		res = res + sep + l
	}
	return res
}
