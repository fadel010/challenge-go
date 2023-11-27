package piscine

func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	sign := 1
	start := 0
	if s[0] == '-' {
		sign = -1
		start = 1
	} else if s[0] == '+' {
		start = 1
	}

	result := 0
	for i := start; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			result = result*10 + int(s[i]-'0')
		} else {
			return 0
		}
	}

	return result * sign
}
