package main

import (
	"github.com/01-edu/z01"
)

func main() {
	result := Rot14("Hello! How are You?")

	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func Rot14(s string) string {
	str := []rune(s)
	for i, v := range str {
		if v >= 'a' && v <= 'z' {
			if v >= 'a' && v <= 'l' {
				str[i] += 14

			} else {
				str[i] -= 12
			}
		} else if v >= 'A' && v <= 'Z' {
			if v >= 'A' && v <= 'L' {
				str[i] += 14

			} else {
				str[i] -= 12
			}
		}
	}
	return string(str)

}
