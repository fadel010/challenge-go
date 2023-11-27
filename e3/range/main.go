package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 2 {
		arg1, err := Atoi(args[0])
		if len(err) != 0 {
			for _, v := range err {
				z01.PrintRune(v)
			}
			z01.PrintRune('\n')

			return
		}
		arg2, err := Atoi(args[1])
		if len(err) != 0 {
			for _, v := range err {
				z01.PrintRune(v)
			}
			z01.PrintRune('\n')

			return
		}
		num := []rune{}
		for i := arg1; i <= arg2; i++ {
			for _, v := range itoa(i) {
				num = append(num, v)
			}
		}
		for i, v := range num {
			z01.PrintRune(v)
			if i != len(num)-1 && v != '-'{
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')

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
			return 0, "strconv.Atoi: parsing " + s + ": invalid syntax"
		}
	}
	return result * sign, ""
}

func itoa(i int) string {

	if i == 0 {
		return "0"
	}
	// determinons le signe deja
	signe := 0
	if i < 0 {
		signe = -1
		i *= -1 // pour pouvoir utiliser i dan sla boucle for sinon la boucle ne va pas s'executer

	} else {
		signe = 1
	}

	// conversion en string
	result := "" //Variable pour stocker le resultat
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
