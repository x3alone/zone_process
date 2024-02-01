package piscine

func SortTable(table *[]int) {
	for i := 0; i < len(*table)-1; i++ {
		for l := i + 1; l < len(*table); l++ {
			if (*table)[i] > (*table)[l] {
				tmp := (*table)[i]
				(*table)[i] = (*table)[l]
				(*table)[l] = tmp
			}
		}
	}
}

func Abort(a, b, c, d, e int) int {
	table := []int{a, b, c, d, e}
	SortTable(&table)
	return table[2]
}
