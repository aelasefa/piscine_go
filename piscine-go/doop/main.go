package main

import "os"

func Atoi(s string) int {
	signe := 1
	r := 0
	for i, char := range s {
		if char == '-' && i == 0 {
			signe *= -1
			continue
		} else if char == '+' && i == 0 {
			continue
		}
		if char >= '0' && char <= '9' {
			r = r*10 + int(char) - '0'
		} else {
			return 0
		}
	}
	return signe * r
}

func IsNumeric(s string) bool {
	if s[0] == '-' {
		s = s[1:]
	}
	for _, char := range s {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}

func PrintNbr(n int) {
	if n < 0 {
		PrintRune('-')
		n *= -1
	}
	if n < 10 {
		PrintRune(rune(n) + '0')
	} else {
		PrintNbr(n / 10)
		PrintRune(rune(n%10) + '0')
	}
}

func PrintStr(s string) {
	for _, v := range s {
		PrintRune(v)
	}
}

func PrintRune(r rune) {
	_, _ = os.Stdout.Write([]byte(string(r)))
}

func main() {
	if len(os.Args) == 4 {
		args := os.Args[1:]

		if IsNumeric(args[0]) && IsNumeric(args[2]) {
			value1 := Atoi(args[0])
			value2 := Atoi(args[2])
			if value1 >= 2147483647 || value2 >= 2147483647 || value1 <= -2147483648 || value2 <= -2147483648 {
				return
			}

			operator := args[1]
			if operator == "+" {
				PrintNbr(value1 + value2)
			} else if operator == "-" {
				PrintNbr(value1 - value2)
			} else if operator == "/" {
				if value2 == 0 {
					PrintStr("No division by 0")
				} else {
					PrintNbr(value1 / value2)
				}
			} else if operator == "*" {
				PrintNbr(value1 * value2)
			} else if operator == "%" {
				if value2 == 0 {
					PrintStr("No modulo by 0")
				} else {
					PrintNbr(value1 % value2)
				}
			} else {
				return
			}
			PrintStr("\n")
		}
	}
}
