package piscine

func Map(f func(int) bool, a []int) []bool {
	c := make([]bool, len(a))
	for i, res := range a {
		c[i] = f(res)
	}
	return c
}
