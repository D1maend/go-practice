package main

import "fmt"

func main() {
	nums := []int{1, 2, 1}
	fmt.Println(getConcatenation(nums))
}
func getConcatenation(nums []int) []int {
	var result []int
	for i := range nums {
		result = append(result, nums[i])
	}

	for i := range nums {
		result = append(result, nums[i])
	}
	return result
}
