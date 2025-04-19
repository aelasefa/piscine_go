package piscine

func ListReverse(l *List) {
	current := l.Head
	var previous *NodeL
	previous = nil
	for current != nil {
		n := current.Next
		current.Next = previous
		previous = current
		current = n
	}
	nn := l.Head
	l.Head = l.Tail
	l.Tail = nn
}
