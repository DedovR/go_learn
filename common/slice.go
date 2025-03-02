package main

import (
	"fmt"
	"slices"
)

func main(){
	a := []int{5, 4, 3, 2, 1}
	fmt.Println(a) // [5 4 3 2 1]
	sort(a)
	fmt.Println(a) // [1 2 3 4 5]
}

func sort(a []int) {
	slices.Sort(a)
}
