package main

import (
	"os"
	"fmt"
)

func main() {
	args := os.args[1:]

	if len(args) != 1 {
		return
	}

	input := Atoi(args[0])
	if input <= 0 || input >= 4000 {
		fmt.Printf("ERROR: cannot convert to roman digit.")
		return
	}

	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	roman1 := []string{"M", "(M-C)", "D", "D-C", "C", "(C-X)", "L", "(L-X)", "X", "(X-I)", "V", "(V-I)", "I"}
	roman2 := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	roman1Result := ""
	roman2Result := ""

	for i := 0; i < len(nums); {
		if nums[i] <= input {
			input -= nums[i]

			roman1Result += roman1[i]
			roman2Result += roman2[i]

			if input != 0 {
				roman1Result += "+"
			}
		} else {
			i++
		}
	}
	fmt.Printf(roman1Result)
	fmt.Printf(roman2Result)

}

func Atoi(s string) int {
	sign := 1
	start := 0

	if s[0] == '-' {
		sign = -1
		start = 1
	} else if s[0] == '+' {
		start = 1
	}

	result := 0

	for i := start; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			result = int(s[i]-'0') + result*10
		} else {
			return 0
		}
	}
	return result * sign
}
