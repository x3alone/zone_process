package piscine

func ListClear(l *List) {
	temp := l.Head
	for temp != nil {
		next := temp.Next
		temp.Next = nil
		temp = next
	}
	l.Head = nil
}
