package piscine

func UltimateDivMod(a *int, b *int) {
	var div int
	var mod int
	div = *a / *b
	mod = *a % *b
	*a = div
	*b = mod
}
