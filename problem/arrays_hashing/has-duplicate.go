package arrays_hashing

func hasDuplicate(nums []int) bool {
	mapNums := make(map[int]bool, len(nums))
	for _, n := range nums {
		if _, existing := mapNums[n]; existing {
			return true
		}
		mapNums[n] = true
	}
	return false
}
