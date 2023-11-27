package main

import (
	"fmt"
	"os"
	"strconv"
)

var romanNumerals = map[int]string{
	1:    "I",
	5:    "V",
	10:   "X",
	50:   "L",
	100:  "C",
	500:  "D",
	1000: "M",
}

func rn(num int) string {
	if num <= 0 || num >= 4000 {
		return "ERROR: cannot convert to roman digit"
	}

	result := ""
	for num > 0 {
		// Find the largest Roman numeral that is less than or equal to num
		var cle int
		var romanValue string
		for v, n := range romanNumerals {
			if num >= v && v > cle {
				cle = v
				romanValue = n
			}
		}

		// Subtract the cle of the Roman romanValue and add it to the result
		result += romanValue
		num -= cle
	}

	return result
}

func main() {
	if len(os.Args[1:]) != 1 {
		return
	}

	numStr := os.Args[1]
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}

	romanNumeral := rn(num)

	fmt.Println(romanNumeral)
}
