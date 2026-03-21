package greedy_heaps

import (
	"sort"
)

func KSmallestElement(arr []int, k int) int {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	return arr[k-1]
}
