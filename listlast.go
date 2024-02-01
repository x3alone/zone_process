package piscine

func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	}
	h := l.Head
	for h.Next != nil {
		h = h.Next
	}
	return h.Data
}
