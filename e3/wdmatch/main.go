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
		empty = append(empty, '2')
		for _, u := range arg1 {
			for _, v := range arg2 {
				if u == v {
					if u != empty[len(empty)-1]{
					empty = append(empty, v)
					}
				}
			}
		}
		empty = empty[1:]
		for _, e := range empty {
			z01.PrintRune(e)
		}
		z01.PrintRune('\n')
	}
}
