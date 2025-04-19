package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	r := max - min
	slice := make([]int, r)
	for i := 0; i < r; i++ {
		slice[i] = i + min
	}
	return slice
}
