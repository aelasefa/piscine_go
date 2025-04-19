package piscine

func AdvancedSortWordArr(table []string, f func(a, b string) int) {
	for i := 0; i < len(table); i++ {
		j := i
		for j > 0 {
			if f(table[j-1], table[j]) == 1 {
				table[j-1], table[j] = table[j], table[j-1]
			}
			j--
		}
	}
}
