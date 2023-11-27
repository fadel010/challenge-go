package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 {
		s := args[0]
		bowl := []rune{}
		cad := []string{}

		for i, v := range s {
			if v != ' ' {
				bowl = append(bowl, v)
			} else {
				cad = append(cad, string(bowl))
				bowl = []rune{}
			}
			if i == len(s)-1 {
				cad = append(cad, string(bowl))

			}

		}

		result := []string{}
		for _, v := range cad {
			if len(v) != 0 {
				result = append(result, v)
			}
		}

		for _, v := range result[1:] {
			for _, w := range v {
				z01.PrintRune(w)
			}
			z01.PrintRune(' ')

		}

		for _, j := range result[0] {
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')

	} else {
		z01.PrintRune('\n')
	}
	

}
