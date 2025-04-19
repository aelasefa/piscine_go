package piscine

import "github.com/01-edu/z01"

func printnbr(n int) {
	z01.PrintRune(rune(n) + '0')
}

func PrintNbrInOrder(n int) {
	var runes []int
	if n <= 0 {
		z01.PrintRune('0')
	}
	for n > 0 {
		runes = append(runes, n%10)
		n /= 10
	}
	for i := 0; i < len(runes)-1; i++ {
		for j := i + 1; j < len(runes); j++ {
			if runes[i] > runes[j] {
				runes[i], runes[j] = runes[j], runes[i]
			} else {
				continue
			}
		}
	}
	for _, pr := range runes {
		printnbr(pr)
	}
}
