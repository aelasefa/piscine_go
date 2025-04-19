package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	y := os.Args[1:]
	for x := 0; x < len(y); x++ {
		for z := x + 1; z < len(y); z++ {
			if y[x] > y[z] {
				y[x], y[z] = y[z], y[x]
			}
		}
		for _, str := range y[x] {
			z01.PrintRune(str)
		}
		z01.PrintRune('\n')
	}
}
