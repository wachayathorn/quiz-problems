package ploblem

func MinimumPlatformsRequired(arr, dep []int) int {
	result := 0

	trains := len(arr)

	for i := 0; i < trains; i++ {
		currentTrainArrTime := arr[i]
		currentTrainDepTime := dep[i]

		for j := i + 1; j < trains; j++ {
			nextTrainArrTime := arr[j]
			nextTrainDepTime := dep[j]

			if currentTrainArrTime >= nextTrainArrTime || currentTrainDepTime <= nextTrainDepTime {
				result++
				continue
			}

			if nextTrainArrTime >= currentTrainArrTime || nextTrainDepTime <= currentTrainDepTime {
				result++
				continue
			}
		}
	}

	return result
}
