package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

var output string = "x = 42, y = 21\n"

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	for _, char := range output {
		z01.PrintRune(char)
	}
}
