package piscine

func ForEach(f func(int), a []int) {
	for _, result := range a {
		f(result)
	}
}
