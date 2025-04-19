package piscine

func Unmatch(a []int) int {
	var q int
	for _, el := range a {
		q = 0
		for _, v := range a {
			if v == el {
				q++
			}
		}
		if q%2 != 0 {
			return el
		}
	}
	return -1
}
