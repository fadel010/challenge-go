package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
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

		for _, v := range str {
			if v == ' ' {
				z01.PrintRune(' ')
			}
			for j, majelemt := range maj {
				// On cherche l'element
				if v == majelemt { // Si l'element est trouve
					k := j + 1

					for k != 0 { // On le print autant de fois
						z01.PrintRune(maj[j])
						k--
					}
				} else if v == min[j] {
					k := j + 1
					for k != 0 { // On le print autant de fois
						z01.PrintRune(min[j])
						k--
					}
				}

				// On cherche les minuscles

			}

		}
	}
	z01.PrintRune('\n')
}
