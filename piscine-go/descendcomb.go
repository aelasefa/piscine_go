package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for i := '9'; i >= '0'; i-- {
		for j := '9'; j >= '0'; j-- {
			d := j - 1
			for c := i; c >= '0'; c-- {
				for ; d >= '0'; d-- {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(' ')
					z01.PrintRune(c)
					z01.PrintRune(d)
					if i > '0' || j > '1' {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
				d = '9'
			}
		}
	}
}
