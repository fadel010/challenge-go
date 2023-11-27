package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 {
		num := Atoi(args[0])
		if num != 0{
			n := 2
			tp := []int{}

			for n <= num {
				if num%n == 0 {
					tp = append(tp, n)
					num /= n
				} else {
					n++
				}
			}
			if len(tp) > 0 {
				for i,v := range tp {
					for _, w := range Itoa(v){
						z01.PrintRune(w)
					}
					if i != len(tp)-1 {
						z01.PrintRune('*')
					}
				}
			}
		
		}
		z01.PrintRune('\n')

	}	
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

func Itoa(i int) string {
	signe := 0
	if i < 0 {
		signe = -1
		i *= -1

	} else {
		signe = 1
	}

	result := ""
	for i > 0 {
		nb := i % 10
		result = string(nb+'0') + result
		i = i / 10
	}
	if signe < 0 {
		result = "-" + result
	}
	return result
}

