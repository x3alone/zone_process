package piscine

func ListSize(l *List) int {
	counter := 1

	noded := l.Head
	for noded.Next != nil {
		noded = noded.Next ; counter++
	}
	return counter 
}