package main

import "fmt"

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}

func Chunk(slice []int, size int) {
	if len(slice) < 1 || size < 1{
		return
	}
	var tab [][]int
	
	for i := 0; i < len(slice);  i += size {


		chucksize := size + i

		if chucksize > len(slice) {
			chucksize = len(slice)
		}
		
		tab = append(tab, slice[i : chucksize])
	}

	fmt.Println(tab)

}