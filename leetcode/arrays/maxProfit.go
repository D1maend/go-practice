package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minPrice := prices[0]
	maxProfit := 0

	for i := 0; i < len(prices); i++ {

		if prices[i] < minPrice {
		} else {
			currentProfit := prices[i] - minPrice

			if currentProfit > maxProfit {
				maxProfit = currentProfit
			}
		}
	}

	return maxProfit
}
