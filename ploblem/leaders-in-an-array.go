package ploblem

func LeadersInAnArray(nums []int) []int {
	var (
		result []int
		sum    int
	)

	for _, num := range nums {
		sum += num
	}

	for i := 0; i < len(nums); i++ {
		currentNum := nums[i]

		suffixMaximum := sum - currentNum

		if suffixMaximum/currentNum == 0 {
			result = append(result, currentNum)
		}

		sum -= currentNum
	}

	return result
}
