package piscine

func LoafOfBread(str string) string {
	if len(str) == 0 {
		return "\n"
	} else if len(str) < 5 {
		return "Invalid Output\n"
	} else {
		res := ""
		i := 0
		for _, a := range str {
			if i%6 != 5 && a == ' ' {
				continue
			}
			if i%6 == 5 {
				res += " "
			} else {
				res += string(a)
			}
			i++
		}
		for i := len(res) - 1; i >= 0; i-- {
			if res[i] != ' ' {
				res = res[:i+1]
				break
			}
		}
		return res + "\n"
	}
}
