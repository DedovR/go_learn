package main

import (
	"fmt"
	"slices"
)

func main(){
	a := []int{5, 4, 3, 2, 1}
	fmt.Println(a)
	sort(a)
	fmt.Println(a)
}

func sort(a []int) {
	slices.Sort(a)
}
