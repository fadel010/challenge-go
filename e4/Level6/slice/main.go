package main

import "fmt"

func main(){
    a := []string{"coding", "algorithm", "ascii", "package", "golang"}
    fmt.Printf("%#v\n", Slice(a, 1))
    fmt.Printf("%#v\n", Slice(a, 2, 4))
    fmt.Printf("%#v\n", Slice(a, -3))
    fmt.Printf("%#v\n", Slice(a, -2, -1))
    fmt.Printf("%#v\n", Slice(a, 2, 0))
}

func Slice(a []string, nbrs... int) []string{
	var tabvide []string
	tab := []int{}
	for _, v := range nbrs {
		tab = append(tab, v)
	}
	for _,v := range tab{
		if v == 0 {
			return tabvide
		}
	}
	if len(tab) == 1 {
		if tab[0] < 0 {
			return a[len(a)+tab[0]:]
		}
		return a[tab[0]:]
	}else if len(tab) == 2 {
		if tab[0] < 0 && tab[1] < 0 {
			return a[len(a)+tab[0]: len(a)+tab[1]]
		}
		return a[tab[0]: tab[1]]
	}
	return a

}