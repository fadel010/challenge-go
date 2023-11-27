package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) == 2 {
		str1 := []rune(args[0])
		str2 := []rune(args[1])
		var result []rune

		for _, u := range str1 {
			for _, v := range str2 {
				if u == v {
					result = append(result, v)
				}
			}
		}

		final := clearResult(result)
		for _, v := range final {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	}
}

func clearResult(result []rune) []rune {
	var fresult []rune
	b := false
	for _, u := range result {
		b = false
		for _, v := range fresult {
			if u == v {
				b = true
			}
		}
		if !b {
			fresult = append(fresult, u)
		}
	}
	return fresult
}
