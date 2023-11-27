package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	} else if len(args) == 1 {

		str := args[0]

		wordList := []string{}

		var bowl []rune

		for i, v := range str {
			if v != ' ' && v != '\t' {
				bowl = append(bowl, v)

			} else {
				wordList = append(wordList, string(bowl))
				bowl = []rune{}
			}

			if i == len(str)-1 {
				wordList = append(wordList, string(bowl))
			}
		}

		result := []string{}

		for _, word := range wordList {
			if len(word) != 0 {

				result = append(result, word)

			}

		}

		for i, word := range result {
			for _, w := range word{
				z01.PrintRune(w)
			}

			if i != len(result) - 1 {
				z01.PrintRune(' ')
			}

		}
	}
	z01.PrintRune('\n')


}
