package piscine

func sp(s string) []string {
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

func Split(s, sep string) []string {
	n := len(s)
	m := len(sep)
	for i := 0; i < n-m; i++ {
		if s[i:i+m] == sep {
			s = s[:i] + " " + s[i+m:]
			n -= m
		}
	}
	return sp(s)
}
