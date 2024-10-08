package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	node := &NodeL{Data: data, Next: nil}
	if l.Head == nil {
		l.Head = node
		return
	}
	noded := l.Head
	for noded.Next != nil {
		noded = noded.Next
	}
	noded.Next = node
}
