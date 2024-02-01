package piscine

func ReverseMenuIndex(menu []string) []string {
	versa := make([]string, len(menu))

	for i := len(menu) - 1; i >= 0; i-- {
		versa[len(menu)-1-i] = menu[i]
	}

	return versa
}
