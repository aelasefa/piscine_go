package piscine

func ListLast(l *List) interface{} {
	if l.Tail == nil && l.Head != nil {
		return l.Head.Data
	} else if l.Tail == nil && l.Head == nil {
		return nil
	}
	return l.Tail.Data
}
