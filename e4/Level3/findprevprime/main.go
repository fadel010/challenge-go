package main

import "fmt"

func main() {
	fmt.Println(FindPrevPrime(5))
	fmt.Println(FindPrevPrime(4))
}

func FindPrevPrime(nb int) int {
	for nb != 2 {
		if IsPrime(nb) {
			return nb
		} else {
			nb--
		}
	}
	return 0
}

func IsPrime(i int) bool {
	d := 2
	for d < i {
		if i%d == 0 {
			return false
		}
		d++
	}
	return true
}
