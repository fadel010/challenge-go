package main

import "fmt"

func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", Split(s, "HA"))
}

func Split(s, sep string) []string {
	n1 := len([]rune(s))
	n2 := len([]rune(sep))

	for i := 0; i < n1-n2; i++ {

		if s[i:n2+i] == sep {
			s = s[:i] + " " + s[n2+i:]
			n1 = n1 - n2
		}

	}
	return SplitWhiteSpaces(s)

}

func SplitWhiteSpaces(s string) []string {
	bowl := []rune{}
	c := []string{}

	for i, u := range s {
		if u != ' ' {
			bowl = append(bowl, u)
		} else {
			c = append(c, string(bowl))
			bowl = []rune{}
		}
		if i == len(s)-1 {
			c = append(c, string(bowl))
		}

	}
	return c
}
