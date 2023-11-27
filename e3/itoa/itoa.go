package main

import "fmt"

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

func main() {
	fmt.Println(itoa(12345))
}
