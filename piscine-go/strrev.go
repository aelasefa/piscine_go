package piscine

func StrRev(s string) string {
	rev := []rune(s)
	i := 0
	j := len(s) - 1
	for i < j {
		rev[i], rev[j] = rev[j], rev[i]
		i++
		j--
	}
	return string(rev)
}
