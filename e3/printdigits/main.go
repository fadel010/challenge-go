package main

import "github.com/01-edu/z01"

func main() {
	for _, d := range "0123456789" {
		z01.PrintRune(d)
	}
	z01.PrintRune('\n')
}
