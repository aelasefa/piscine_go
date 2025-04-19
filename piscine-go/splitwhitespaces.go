package piscine

func SplitWhiteSpaces(s string) []string {
	var slice []string
	r := ""
	for _, char := range s {
		if char != ' ' {
			r += string(char)
		} else if r != "" {
			slice = append(slice, r)
			r = ""
		}
	}
	if r != "" {
		slice = append(slice, r)
	}
	return slice
}
