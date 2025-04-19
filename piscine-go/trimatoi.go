package piscine

func TrimAtoi(s string) int {
	var r string
	for _, char := range s {
		if char >= '0' && char <= '9' || char == '-' {
			r += string(char)
		}
	}
	return atoi(r)
}

func atoi(s string) int {
	r := 0
	signe := 1
	for i, char := range s {
		if char == '-' && i == 0 {
			signe *= -1
			continue
		} else if char == '-' && i > 0 {
			continue
		} else {
			r = r*10 + int(char) - '0'
		}
	}
	return signe * r
}
