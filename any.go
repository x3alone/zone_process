package piscine

func Any(f func(string) bool, a []string) bool {
	is := false
	for _, i := range a {
		if f(i) {
			return true
		}
	}
	return is
}
