package main

import (
	"os"

	"github.com/01-edu/z01"
)

func masin() {
	args := os.Args[1:]
	if len(args) == 1 {

		str := []rune(args[0])
		var maj []rune
		var min []rune

		for i := 65; i <= 90; i++ {
			maj = append(maj, rune(i))
		}
		for i := 97; i <= 122; i++ {
			min = append(min, rune(i))
		}

		for i, v := range str {
			for j, majelemt := range maj {
				if v == majelemt {
					str[i] = min[j]
				} else if v == min[j] {
					str[i] = maj[j]
				}
			}

		}

		for _, v := range str {
			z01.PrintRune(v)

		}
	}
	z01.PrintRune('\n')

}
