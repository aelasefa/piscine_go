package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	arg := os.Args[1:]
	if len(arg) < 1 {
		fmt.Println("File name missing")
		return
	}
	if len(arg) > 1 {
		fmt.Println("Too many arguments")
		return
	}
	variable, _ := ioutil.ReadFile(arg[0])
	s := ""
	for _, char := range variable {
		if char >= ' ' && char <= '~' {
			s = s + string(rune(char))
		}
	}
	fmt.Println(s)
}
