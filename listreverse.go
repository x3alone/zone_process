package piscine

func ListReverse(l *List) {
	temp := l.Head
	temp2 := l.Head
	temp2 = nil
	for temp != nil {
		next := temp.Next
		temp.Next = temp2
		temp2 = temp
		temp = next
	}
	s := l.Head
	l.Head = l.Tail
	l.Tail = s
}
