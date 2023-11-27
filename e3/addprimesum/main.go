package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 {
		n := Atoi(args[0])
		if n > 1 {
			tab1 := []int{}
			tab2 := []int{}

			for i := 2; i <= n; i++ {
				tab1 = append(tab1, i)
			}

			for _, v := range tab1 {
				if isPrime(v) {
					tab2 = append(tab2, v)
				}
			}
			var primesum int

			for _, v := range tab2 {
				primesum += v
			}
			for _, v := range itoa(primesum) {
				z01.PrintRune(v)
			}
			z01.PrintRune('\n')
		} else {
			z01.PrintRune('0')
			z01.PrintRune('\n')
			return
		}
	} else {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

}

func isPrime(i int) bool {
	if i < 0 {
		return false
	}
	for j := 2; j < i; j++ {
		if i%j == 0 {
			return false
		}
	}

	return true
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
