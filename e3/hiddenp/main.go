package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 2 {
		arg1 := []rune(args[0])
		arg2 := []rune(args[1])
		var empty []rune
		b := false

		for _, v := range arg1 {
			b = false
			for _, w := range arg2 {
				if v == w {

					if len(empty) > 0 && w == empty[len(empty)-1] {
						b = true

					}

					if !b {
						empty = append(empty, w)
					}
				}
			}
		}
		if string(empty) == string(arg1) {
			z01.PrintRune('1')
			z01.PrintRune('\n')

		} else {
			z01.PrintRune('0')
			z01.PrintRune('\n')

		}
	}
}
