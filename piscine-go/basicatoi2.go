package piscine

func BasicAtoi2(s string) int {
	r := 0
	for _, char := range s {
		if char >= '0' && char <= '9' {
			r = r*10 + int(char) - '0'
		} else {
			return 0
		}
	}
	return r
}
