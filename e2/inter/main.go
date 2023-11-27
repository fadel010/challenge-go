package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 2 {
		str1 := []rune(args[0])
		str2 := []rune(args[1])
		var result []rune

		for _, u := range str1 {
			for _, v := range str2 {
				if u == v {
					result = append(result, v)
				}
			}
		}

		final := clearResult(result)
		fmt.Println(final)
	}
}

func clearResult(result []rune) string{
	var fresult []rune
	b := false
	for _, u := range result {
		b = false
		for _, v := range fresult {
			if u == v {
				b = true
			}
		}
		if !b {
			fresult = append(fresult, u)
		}
	}
	return string(fresult)
}