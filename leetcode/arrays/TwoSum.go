package main

import "fmt"

func main() {
	twonum := [4]int{2, 11, 7, 15}
	target := 9

	fmt.Println(twoSum(twonum, target))
}
func twoSum(nums [4]int, target int) [2]int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {

			if nums[i]+nums[j] == target {

				return [2]int{i, j}
			}

		}

	}
	return [2]int{0, 0}

}
