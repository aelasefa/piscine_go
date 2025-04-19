package piscine

func ToLower(s string) string {
	var r string
	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			r += string(char + 32)
		} else {
			r += string(char)
		}
	}
	return r
}
