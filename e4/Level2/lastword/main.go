package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 {
		str := []rune(args[0])
		var lastword []rune
		for i := len(str) - 1; i >= 0; i-- {
			if str[i] == ' ' {
				break
			}
			lastword = append(lastword, str[i])
		}
		for i := len(lastword) - 1; i >= 0; i-- {
			z01.PrintRune(lastword[i])
		}
		z01.PrintRune('\n')

	}

}
