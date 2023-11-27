package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	} else {
		for _, v := range args {

			GoodList := ReverseCap(CreateWordList(v))

			for i, word := range GoodList {
				for _, f := range word {
					z01.PrintRune(f)
				}
				if i != len(GoodList)-1 {
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune('\n')

		}
	}
}

func CreateWordList(str string) []string {
	wordList := []string{}

	var bowl []rune

	for i, v := range str {
		if v != ' ' {
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

	return result
}

func ReverseCap(word []string) []string {
	runeTab := []rune{}

	for j, v := range word {
		runeTab = []rune(v)

		for i, w := range runeTab {
			if i == len(v)-1 {
				if w >= 'a' && w <= 'z' {
					runeTab[i] -= 32
				}
			} else {
				if w >= 'A' && w <= 'Z' {
					runeTab[i] += 32
				}
			}
		}
		word[j] = string(runeTab)

	}
	return word
}
