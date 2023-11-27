 package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 {
		r := []rune(args[0])
		var empty []rune
		for i, u := range r {
			if u != ' ' {
				empty = append(empty, u)
			} else if u == ' ' {
				if i != 0 && i != len(r)-1 {
					k := 3
					for k != 0 {
						empty = append(empty, ' ')
						k--
					}
				}
			}
		}
		for _, v := range empty {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	}
}
