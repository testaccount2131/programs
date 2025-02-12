package main

import "fmt"

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
func bubbleSortrecur(arr []int, n int) {
	if n <= 1 {
		return
	}
	for i := 0; i < n-1; i++ {
		if arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}

	}
	bubbleSortrecur(arr, n-1)
}
func selectionsort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}
func recursiveSelectionSort(arr []int, start int) {
	n := len(arr)
	// Base case: if the array has 0 or 1 element, it's already sorted
	if start >= n-1 {
		return
	}

	// Find the index of the minimum element in the unsorted portion
	minIndex := start
	for i := start + 1; i < n; i++ {
		if arr[i] < arr[minIndex] {
			minIndex = i
		}
	}

	// Swap the minimum element with the first unsorted element
	arr[start], arr[minIndex] = arr[minIndex], arr[start]

	// Recursively sort the remaining unsorted portion of the array
	recursiveSelectionSort(arr, start+1)
}

func insertionsort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--

		}
		arr[j+1] = key
	}
}
func insertionsortrecur(arr []int, index int) {
	n := len(arr)
	if index == n {
		return
	}

	j := index - 1
	key := arr[index]
	for j >= 0 && arr[j] > key {
		arr[j+1] = arr[j]
		j--
	}
	arr[j+1] = key

	insertionsortrecur(arr, index+1)
}

func quicksort(arr []int) []int {
	n := len(arr)
	if n <= 1 { // Check for base case first
		return arr
	}
	pivot := arr[0] // Now it's safe to assign the pivot
	var left, right []int
	for i := 1; i < n; i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}
	left = quicksort(left)
	right = quicksort(right)
	return append(append(left, pivot), right...)
}

func mergesort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}
	mid := n / 2
	leftsort := mergesort(arr[:mid])
	rightsort := mergesort(arr[mid:])
	return merge(leftsort, rightsort)
}
func merge(arr1 []int, arr2 []int) []int {
	total := make([]int, 0, len(arr1)+len(arr2))
	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			total = append(total, arr1[i])
			i++
		} else {
			total = append(total, arr2[j])
			j++
		}
	}
	total = append(total, arr1[i:]...)
	total = append(total, arr2[j:]...)
	return total
}
func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	arr1 := mergesort(arr)
	fmt.Println("Sorted array:", arr1)
}
