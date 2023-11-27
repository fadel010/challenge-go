package main

import "fmt"

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}

	for i := 0; i < len(base); i++ {
		if base[i] == '+' || base[i] == '-' {
			return false
		}
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				return false
			}
		}
	}

	return true
}

func charToDigit(ch byte) int {
	if ch >= '0' && ch <= '9' {
		return int(ch - '0')
	} else if ch >= 'A' && ch <= 'Z' {
		return int(ch - 'A' + 10)
	} else if ch >= 'a' && ch <= 'z' {
		return int(ch - 'a' + 10)
	}
	return -1
}

func AtoiBase(s string, base string) int {
	if !isValidBase(base) {
		return 0
	}

	baseLen := len(base)
	result := 0

	for i := 0; i < len(s); i++ {
		digitValue := charToDigit(s[i])
		if digitValue == -1 || digitValue >= baseLen {
			return 0
		}
		result = result*baseLen + digitValue
	}

	return result
}

func main() {
	fmt.Println(AtoiBase("125", "0123456789"))
	fmt.Println(AtoiBase("1111101", "01"))
	fmt.Println(AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(AtoiBase("uoi", "choumi"))
	fmt.Println(AtoiBase("bbbbbab", "-ab"))
}
