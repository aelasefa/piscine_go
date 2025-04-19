package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	if l == nil {
		return nil
	}
	ListSlice := []NodeL{}
	current := l
	for current != nil {
		if current.Next == nil {
			ListSlice = append(ListSlice, *current)
			current = nil
		} else {
			ListSlice = append(ListSlice, *current)
			current = current.Next
		}
	}
	if pos > len(ListSlice)-1 || pos < 0 {
		return nil
	}
	return &ListSlice[pos]
}
