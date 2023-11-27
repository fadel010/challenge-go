package main

import (
	"fmt"
)

func main() {
	a := []int{23, 13, 1, 11, 55, 93}
	max := Max(a)
	fmt.Println(max)
}

func Max(a []int) int {
	max := a[0]
	for _, d := range a {
		if max < d {
			max = d
		}

	}
	return max

}
