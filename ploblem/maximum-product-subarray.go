package ploblem

import (
	"fmt"
)

func MaxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var (
		sum        int
		numsResult []int
	)

	for i, num := range nums {
		currentNumsResult := []int{num}
		currentValue := num
		for j := i + 1; j < len(nums); j++ {
			currentNumsResult = append(currentNumsResult, nums[j])
			currentValue = currentValue * nums[j]
			if sum < currentValue {
				numsResult = currentNumsResult
				sum = currentValue
			}
			if sum > currentValue && j != i+1 {
				break
			}
		}
	}

	fmt.Println(sum)
	fmt.Println(numsResult)

	return sum
}
