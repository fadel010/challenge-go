package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	num, _ := strconv.Atoi(args[0])

	if num > 0 {
		n := '1'
		m := 1
		for n <= '9' {
			z01.PrintRune(n)
			z01.PrintRune(' ')
			z01.PrintRune('*')
			z01.PrintRune(' ')

			for _, v := range args[0] {
				z01.PrintRune(v)
			}
			z01.PrintRune(' ')
			z01.PrintRune('=')
			z01.PrintRune(' ')
			char := strconv.Itoa(m * num)
			for _, c := range char {
				z01.PrintRune(c)
			}
			z01.PrintRune('\n')
			m++
			n++

		}
	}
}
