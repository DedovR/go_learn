package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(uniqRandn(10))
}

func uniqRandn(n int) []int {
	arr := make([]int, n)
	uMap := map[int]bool{}
	i := 0
	for i < n {
		randI := rand.Intn(15)
		_, ok := uMap[randI]
		fmt.Println(randI)
		if !ok {
				arr[i] = randI
				uMap[randI] = true
				i++
		}
	}

	return arr
}
