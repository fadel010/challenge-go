package main

import (
	"fmt"
)

func main() {
	fmt.Println(Capitalize("Hello! How are you? How+are+things+4you?"))
}

func Capitalize(s string) string {
	char := []rune(s)
	compt := 0
	for i := range char {
		if (char[i] >= 'a' && char[i] <= 'z') || (char[i] >= 'A' && char[i] <= 'Z') || (char[i] >= '0' && char[i] <= '9') {
			compt++
		} else {
			compt = 0
		}
		if compt == 1 && (char[i] >= 'a' && char[i] <= 'z') {
			char[i] -= 32
		} else if compt > 1 && (char[i] >= 'A' && char[i] <= 'Z') {
			char[i] += 32
		}
	}
	return string(char)
}
