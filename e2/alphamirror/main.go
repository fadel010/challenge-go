package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	tabrune := []rune(args[0])
	if len(args) == 1 {
		for _, v := range tabrune {
			if v >= 'a' && v <= 'z' {
				z01.PrintRune('z' - (v - 'a'))

			} else if v >= 'A' && v <= 'Z' {
				z01.PrintRune('Z' - (v - 'A'))

			} else {
				z01.PrintRune(v)
			}
		}
	}
	z01.PrintRune('\n')

}
