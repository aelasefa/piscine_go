package piscine

func ConcatParams(args []string) string {
	var r string
	for i := 0; i < len(args); i++ {
		r += args[i]
		if i != len(args)-1 {
			r += "\n"
		}
	}
	return r
}
