package piscine

func Join(strs []string, sep string) string {
	r := ""
	for i, char := range strs {
		r += char
		if i < len(strs)-1 {
			r += sep
		}
	}
	return r
}
