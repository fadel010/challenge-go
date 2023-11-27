package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 2 {
		r := []rune(args[0] + args[1])
		var tabVide []rune
		find := false
		for _, u := range r {
			find = false
			for _, v := range tabVide {
				if u == v {
					find = true
				}
			}

			if !find {
				tabVide = append(tabVide, u)
			}

		}

		for _, v := range tabVide {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	} else {
		z01.PrintRune('\n')
	}
}
