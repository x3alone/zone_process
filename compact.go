package piscine

func Compact(ptr *[]string) int {
	c := 0
	var array []string
	for _, val := range *ptr {
		if val != "" {
			array = append(array, val)
			c++
		}
	}
	*ptr = array
	return c
}
