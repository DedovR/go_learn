package main

import "math"

func main() {
	arr := []int{1, 2, 3, 4, 5, 7, 8}

	println(BinarySearch(arr, 4))
}

func BinarySearch[T ~int | ~float64](arr []T, value T) int {
	left := 0
	right := len(arr) - 1
	var mid int

	for left <= right {
		mid = int(math.Floor(float64((left + right) / 2)))

		if arr[mid] == value {
			return mid
		}
		if arr[mid] > value {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}
