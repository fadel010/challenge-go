package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 2 {
		arg1 := Atoi(args[0])
		arg2 := Atoi(args[1])
		if arg1 > 0 && arg2 > 0 {
			d := gcd(arg1, arg2)
			fmt.Println(d)
		}
	}
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func Atoi(s string) int {
	sign := 1
	start := 0

	if s[0] == '-' {
		sign *= -1
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
