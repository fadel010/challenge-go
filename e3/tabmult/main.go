package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if Atoi(args[0]) > 0 {
		n := '1'
		m := 1
		for n <= '9' {
			z01.PrintRune(n)
			z01.PrintRune(' ')
			z01.PrintRune('*')
			z01.PrintRune(' ')

			for _, v := range args[0] {
				z01.PrintRune(v)
			}
			z01.PrintRune(' ')
			z01.PrintRune('=')
			z01.PrintRune(' ')
			for _, c := range itoa(m * Atoi(args[0])) {
				z01.PrintRune(c)
			}
			z01.PrintRune('\n')
			m++
			n = n + 1

		}
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

func itoa(i int) string {
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
