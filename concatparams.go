package piscine

func ConcatParams(args []string) string {
	var res string
	for i := 0; i < len(args); i++ {
		res += args[i]
		if i < len(args)-1 {
			res += "\n"
		}
	}
	return res
}
