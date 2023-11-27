package main

import "fmt"

func GetCharIndex(ch rune, base []rune) int {
	for ind, val := range base {
		if val == ch {
			return ind
		}
	}
	return -1
}

func AtoiBase(s string, base string) int {

	for _, val := range base {
		if val == '-' || val == '+' {
			return 0
		}

	}
	if len(base) < 2 {
		return 0
	}
	for i := 0; i < len(base)-1; i++ {
		for j := i + 1; j < len(base); j++ {
			if base[j] == base[i] {
				return 0

			}
		}
	}

	sRune := []rune(s)
	baseRune := []rune(base)


	baseNumb := 1
	res := 0

	for i := 1; i < len(s); i++ {
		baseNumb *= len(baseRune)
	}

	for _, v := range sRune{
		res += baseNumb *  GetCharIndex(v, baseRune)
		baseNumb /= len(baseRune)
	}
	return res
}

func main() {
	fmt.Println(AtoiBase("125", "0123456789"))
	fmt.Println(AtoiBase("1111101", "01"))
	fmt.Println(AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(AtoiBase("uoi", "choumi"))
	fmt.Println(AtoiBase("bbbbbab", "-ab"))
}
