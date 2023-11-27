package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 {
		word := []rune(args[0])
		vowels := []rune{'a', 'e', 'i', 'o', 'u', 'y'}
		bowl := []rune{}

		// Check if the first element of arg is a vowel
		for _, v := range vowels {
			if v == word[0] {
				w := args[0] + "ay"
				for _, v := range w {
					z01.PrintRune(v)
				}
				z01.PrintRune('\n')
				return
			}
		}

		vowelFound := false

		for _, v := range word {
			for _, w := range vowels {
				if v == w {
					vowelFound = true
				}
			}
			if !vowelFound {
				bowl = append(bowl, v)
			}
		}

		if len(bowl) == len(args[0]) {

			for _, v := range "No vowels" {
				z01.PrintRune(v)
			}
			z01.PrintRune('\n')
		} else {
			w := string(word[len(bowl):]) + string(bowl) + "ay"
			for _, v := range w {
				z01.PrintRune(v)
			}
			z01.PrintRune('\n')
		}
	}
}
