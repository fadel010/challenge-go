package main

import "fmt"

func ReduceInt(a []int, f func(int, int) int) {
	result := 0
	for i := 0; i < len(a)-1; i++ {
		result += f(a[i], a[i+1])
	}
	fmt.Println(result)
}

func sum(a, b int) int {
	return a + b
}

func main() {
	a := []int{500, 2}
	ReduceInt(a, sum)
}
