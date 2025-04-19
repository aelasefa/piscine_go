package piscine

func ReverseMenuIndex(menu []string) []string {
	mLen := len(menu)
	r := make([]string, mLen)

	for i, n := range menu {
		NewI := mLen - i - 1
		r[NewI] = n
	}
	return r
}
