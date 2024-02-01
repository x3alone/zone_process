package piscine

func StrRev(s string) string {
	var rev string = ""
	for _, char := range s {
		rev = string(char) + rev
	}
	return rev
}
