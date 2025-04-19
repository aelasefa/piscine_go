package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	head := l

	if l == nil || l.Next == nil {
		return l
	}

	for head != nil {
		if head.Next == nil {
			break
		}
		head = head.Next
	}

	head.Next = &NodeI{
		Data: data_ref,
	}

	current := l

	for current != nil {
		next := current.Next
		for next != nil {
			if current.Data > next.Data {
				current.Data, next.Data = next.Data, current.Data
			}
			next = next.Next
		}
		current = current.Next
	}

	return l
}
