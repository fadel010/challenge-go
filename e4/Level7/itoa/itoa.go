package piscine

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
