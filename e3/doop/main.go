package main

import (
	"os"
)

func main() {

	args := os.Args[1:]
	if len(args) == 3 {
		a, err1 := Atoi(args[0])
		sign := args[1]
		b, err2 := Atoi(args[2])

		if len(err1) != 0 || len(err2) != 0 {
			return
		}

		if a > 9223372036854775806 || b > 9223372036854775806 || a < -9223372036854775806 || b < -9223372036854775806 {
			return
		}

		result := 0

		switch sign {
		case "+":
			result = a + b
		case "-":
			result = a - b
		case "*":
			result = a * b
		case "/":
			if b != 0 {
				result = a / b
			} else {
				strToByte("No division by 0")
				strToByte("\n")

				return
			}
		case "%":
			if b != 0 {
				result = a % b
			} else {
				strToByte("No modulo by 0")
				strToByte("\n")
				return
			}
		default:
			return
		}
		intToBye(result)

		// Check if a and b are valid numbers

	}
}

func strToByte(s string) {
	tab := []byte{}
	for i := range s {
		tab = append(tab, s[i])
	}
	os.Stdout.Write(tab)
}

func Atoi(s string) (int, string) {
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
			return 0, "error"
		}
	}
	return result * sign, ""
}

func intToBye(i int) {
	intTostr := Itoa(i)
	strToByte(intTostr)

}

func Itoa(i int) string {
	if i == 0 {
		return "0"
	}
	signe := 0
	if i < 0 {
		signe = -1
		i = -i

	} else {
		signe = 1
	}
	var result string
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
