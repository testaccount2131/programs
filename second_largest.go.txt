package main

import (
	"fmt"
	"sort"
)

func second(a []int) int {
	sort.Ints(a)
	for j := len(a) - 1; j >= 0; j-- {
		if a[j] != a[len(a)-1] {
			return a[j]
		}
	}
	return 0
}

func main() {
	v := []int{1, 5, 6, 6, 2}
	number := second(v)
	fmt.Println("Second Largest Element is:", number)
}
