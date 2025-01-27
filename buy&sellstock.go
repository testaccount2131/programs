package main // Easy Question//

import (
	"fmt"
)

func maxprofit(a []int) int {
	if len(a) == 0 {
		return 0
	}
	min := a[0]
	max := 0
	for _, price := range a { // Use price instead of index
		if price < min { // Compare price with min
			min = price
		} else {
			profit := price - min
			if profit > max {
				max = profit
			}
		}
	}
	return max
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	profit := maxprofit(prices)
	fmt.Println("Maximum Profit:", profit)
}
