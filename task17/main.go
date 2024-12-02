package main

import "fmt"

func BinarySearch(arr []int, val int) int {
	left, right := 0, len(arr)
	mid := (left + right) / 2

	for left < right {
		if arr[mid] == val {
			return mid
		} else if arr[mid] < val {
			left = mid
		} else {
			right = mid
		}
		mid = (left + right) / 2
	}
	return -1
}

func main() {

	var sortedArray = []int{2, 4, 6, 9, 11, 33, 52, 60, 134, 204, 289, 321, 453}
	var searchNumber = 44
	fmt.Println("Search Number Index: ", BinarySearch(sortedArray, searchNumber))
}