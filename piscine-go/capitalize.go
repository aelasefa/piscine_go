package piscine

func toupper(char rune) rune {
	if char >= 'a' && char <= 'z' {
		char -= 32
	}
	return char
}

func toulower(char rune) rune {
	if char >= 'A' && char <= 'Z' {
		char += 32
	}
	return char
}

func Capitalize(s string) string {
	runes := []rune(s)
	next := true
	for i, char := range s {
		if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' {
			if next {
				runes[i] = toupper(char)
				next = false
			} else {
				runes[i] = toulower(char)
			}
		} else {
			next = true
		}
	}
	return string(runes)
}
