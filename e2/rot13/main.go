package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 {
		str := []rune(args[0])
		for i, v := range str {
			if v >= 'a' && v <= 'z' {
				if v >= 'a' && v <= 'm' {
					str[i] += 13

				} else {
					str[i] -= 13
				}
			} else if v >= 'A' && v <= 'Z' {
				if v >= 'A' && v <= 'M' {
					str[i] += 13

				} else {
					str[i] -= 13
				}
			}
		}
		for _, v := range str {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')

	}

}
