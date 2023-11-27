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
		tabVide = append(tabVide, '2')
		find := false
		for _, u := range r {
			for _, v := range tabVide {
				if u == v {
					find = true
				}
			}

			if !find {
				tabVide = append(tabVide, u)

			}

		}

		for i, v := range tabVide {
			if i != 0 {
				z01.PrintRune(v)
			}
		}
		z01.PrintRune('\n')
	}
}
