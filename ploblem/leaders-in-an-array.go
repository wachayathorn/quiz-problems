package ploblem

func LeadersInAnArray(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}

	result := []int{}
	maxRight := nums[n-1]
	result = append(result, maxRight)

	for i := n - 2; i >= 0; i-- {
		if nums[i] >= maxRight {
			maxRight = nums[i]
			result = append(result, maxRight)
		}
	}

	// Reverse the result slice to maintain original order
	var resultReversed []int = make([]int, len(result))
	for i, v := range result {
		resultReversed[len(result)-1-i] = v
	}

	return resultReversed
}
