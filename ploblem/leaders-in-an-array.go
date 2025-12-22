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
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}
