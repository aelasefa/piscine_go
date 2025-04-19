package piscine

func Swap(a *int, b *int) {
	var t int
	t = *b
	*b = *a
	*a = t
}
