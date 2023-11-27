package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 {
		str := args[0]
		bowl := []rune{}
		thine := []string{}

		for i, v := range str {
			if v != ' ' {
				bowl = append(bowl, v)
			} else {
				thine = append(thine, string(bowl))
				bowl = []rune{}
			}
			if i == len(str)-1 {
				thine = append(thine, string(bowl))
			}

		}
		result := []string{}
		for _, v := range thine {
			if len(v) != 0 {
				result = append(result, v)
			}
		}

		for i := len(result) - 1; i >= 0; i-- {
			for _, v := range result[i] {
				z01.PrintRune(v)
			}
			if i != 0 {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}

}
