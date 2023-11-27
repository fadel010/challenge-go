package main

import "fmt"


func main() {
	fmt.Println(Itoa(11))

}

func Atoi(s string) (int,string) {
	sign := 1
	start := 0

	if s[0] == '-' {
		sign = -1
		start = 1
	} else if s[0] == '+' {
		start = 1
	}
	result := 0

	for i := start; i < len(s); i++  {
		if s[i] >= '0' && s[i] <= '9' {
			result = int(s[i]-'0') + result*10
		} 	else {
			return 0, "error"
		}
	}
	return result*sign,""

}

func Itoa(i int) string{
	if i == 0 {
		return "0"
	}

	sign := 0

	if i < 0 {
		sign = -1
		i = -i
	}
	result := ""

	for i != 0 {
		result = string(i % 10 + '0') + result
		i /= 10
	}

	if sign < 0 {
		result = "-" + result
	}
	return result
}