package ploblem

func MaximumLengthBitonicSubarray(arr []int) int {
	length := len(arr)
	if length == 0 {
		return 0
	}

	// inc[i] = length of increasing sequence ending at index i
	inc := make([]int, length)
	inc[0] = 1
	for i := 1; i < length; i++ {
		if arr[i] >= arr[i-1] {
			inc[i] = inc[i-1] + 1
		} else {
			inc[i] = 1
		}
	}

	// dec[i] = length of decreasing sequence starting at index i
	dec := make([]int, length)
	dec[length-1] = 1
	for i := length - 2; i >= 0; i-- {
		if arr[i] >= arr[i+1] {
			dec[i] = dec[i+1] + 1
		} else {
			dec[i] = 1
		}
	}

	// Find maximum bitonic length
	maxLen := 1
	for i := 0; i < length; i++ {
		bitonicLen := inc[i] + dec[i] - 1
		if bitonicLen > maxLen {
			maxLen = bitonicLen
		}
	}

	return maxLen
}
