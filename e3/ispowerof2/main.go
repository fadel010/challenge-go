package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 {
		num := Atoi(args[0])

		if num > 0 {
			if isPowerof2(num) {
				for _, v := range "true" {
					z01.PrintRune(v)
				}
				z01.PrintRune('\n')
				return
			} else {
				for _, v := range "false" {
					z01.PrintRune(v)
				}
				z01.PrintRune('\n')
				return

			}
		}
	}

}

func isPowerof2(nombre int) bool {
	return nombre&(nombre-1) == 0
}

func Atoi(s string) int {
	sign := 1
	start := 0

	if s[0] == '-' {
		sign *= -1
		start = 1
	} else if s[0] == '+' {
		start = 1
	}
	result := 0

	for i := start; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			result = int(s[i]-'0') + result*10
		} else {
			return 0
		}
	}
	return result * sign
}
