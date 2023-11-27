package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		z01.PrintRune('0')
	} else {
		l := Itoa(len(args))
		for _, v := range l {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	}
}

func Itoa(i int) string {
	signe := 1
	if i < 0 {
		signe = -1
		i = -i
	}

	var result string
	for i > 0 {
		result = string(i%10+'0') + result
		i = i / 10
	}
	if signe < 0 {
		result = "-" + result
	}
	return result
}
