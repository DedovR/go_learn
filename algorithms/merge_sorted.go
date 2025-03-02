package main

import "fmt"

func main() {
	a1 := []int{1,2,6}
	a2 := []int{3,5,21,67}

	a := MergeSorted(a1, a2)
	fmt.Println(a)
}

func MergeSorted[T ~int | ~float64](arr1, arr2 []T) []T {
	arrLen := len(arr1) + len(arr2)
	merged := make([]T, arrLen)
	var i, i1, i2 int
	len1 := len(arr1)
	len2 := len(arr2)

	for i1 < len1 || i2 < len2 {
		if i1 == len1 || arr1[i1] > arr2[i2]{
			merged[i] = arr2[i2]
			i2++
		} else {
			merged[i] = arr1[i1]
			i1++
		}
		i++
	}

	return merged
}
