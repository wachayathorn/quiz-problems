package ploblem

func EquilibriumIndex(nums []int) int {
	for i := range nums {
		sumLeft := 0
		if i > 0 {
			leftSide := nums[:i]
			for _, num := range leftSide {
				sumLeft += num
			}
		}

		rightSide := nums[i+1:]
		sumRight := 0
		for _, num := range rightSide {
			sumRight += num
		}

		if sumLeft == sumRight {
			return i
		}
	}
	return -1
}
