package ploblem

func MinimumPlatformsRequired(arr, dep []int) int {
	trains := len(arr)

	result := 1

	for i := 0; i < trains; i++ {
		currentTrainArrTime := arr[i]
		currentTrainDepTime := dep[i]

		platform := 1

		for j := i + 1; j < trains; j++ {
			nextTrainArrTime := arr[j]
			nextTrainDepTime := dep[j]

			if currentTrainArrTime >= nextTrainArrTime && currentTrainDepTime <= nextTrainDepTime {
				platform++
				continue
			}

			if nextTrainArrTime <= currentTrainDepTime {
				platform++
				continue
			}
		}

		if result < platform {
			result = platform
		}
	}

	return result
}
