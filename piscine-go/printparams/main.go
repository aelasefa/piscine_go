package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	input := os.Args[1:]
	for _, args := range input {
		for _, char := range args {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
