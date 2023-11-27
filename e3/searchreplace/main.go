package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 3 {
		arg1 := []rune(args[0])
		arg2 := []rune(args[1])
		arg3 := []rune(args[2])
		for i, u := range arg1 {
			for _, v := range arg2 {
				if u == v {
					arg1[i] = arg3[0]

				}
			}
		}
		for _, v := range arg1 {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	}

}
