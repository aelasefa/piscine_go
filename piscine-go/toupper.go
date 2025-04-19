package piscine

func ToUpper(s string) string {
	var r string
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			r += string(char - 32)
		} else {
			r += string(char)
		}
	}
	return r
}
