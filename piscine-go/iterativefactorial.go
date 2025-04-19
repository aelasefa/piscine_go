package piscine

func IterativeFactorial(nb int) int {
	var r int
	if nb > 20 {
		r = 0
	} else if nb == 0 || nb == 1 {
		r = 1
	} else if nb > 1 {
		r = 1
		for i := 2; i <= nb; i++ {
			r *= i
		}
	} else {
		return 0
	}
	return r
}
