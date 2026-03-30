package arrays_hashing

func TwoSum(nums []int, target int) []int {
	result := []int{}
	mapNums := make(map[int]int, len(nums))

	for i, n := range nums {
		mapNums[n] = i
	}

	for i, n := range nums {
		diff := target - n
		diffIndex, found := mapNums[diff]
		if found && i != diffIndex {
			result = append(result, i, diffIndex)
			return result
		}
	}

	return result
}
