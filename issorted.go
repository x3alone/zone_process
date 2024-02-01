package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	as, ds := true, true
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) < 0 {
			ds = false
		}
	}
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			as = false
		}
	}
	return as || ds
}
