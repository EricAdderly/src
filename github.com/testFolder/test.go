package main

import (
	"fmt"
)

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 12
	fmt.Println(search(nums, target))
}

func search(nums []int, target int) int {
	var array []int
	array = nums
	localTarget := target
	low := 0
	high := len(array) - 1

	for low <= high {
		median := (low + high) / 2

		if array[median] < localTarget {
			low = median + 1
		} else if array[median] > localTarget {
			high = median - 1
		} else {
			return median
		}
	}
	return -1
}
