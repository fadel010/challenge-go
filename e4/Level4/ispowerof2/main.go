package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 {
		num, _ := strconv.Atoi(args[0])

		if num > 0 {
			if isPowerof2(num) {
				for _, v := range "true" {
					z01.PrintRune(v)
				}
				z01.PrintRune('\n')
				return
			} else {
				for _, v := range "false" {
					z01.PrintRune(v)
				}
				z01.PrintRune('\n')
				return

			}
		}
	}

}

func isPowerof2(nombre int) bool {
	return nombre&(nombre-1) == 0
}
