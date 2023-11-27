package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 {
		num, err := strconv.Atoi(args[0])
		if err != nil {
			for _, v := range "ERROR" {
				z01.PrintRune(v)
			}
			z01.PrintRune('\n')
		}
		hexdecimal := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}
		var result string

		for num > 0 {
			result = hexdecimal[num%16] + result
			num /= 16
		}
		for _, v := range result {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	}
}
