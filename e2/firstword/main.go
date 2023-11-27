package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 {
		str := []rune(args[0])
		var firstword []rune
		for _, v := range str {
			if v == ' ' {
				break
			}
			firstword = append(firstword, v)
		}
		for _, v := range firstword {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')

	}

}
