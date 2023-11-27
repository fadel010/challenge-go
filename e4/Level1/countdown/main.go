package main

import "github.com/01-edu/z01"

func main() {
	for _, v := range "9876543210" {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
