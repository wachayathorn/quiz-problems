package ploblem

import "fmt"

func TrapRainWater(nums []int) int {
	width := len(nums)
	height := 0

	for _, num := range nums {
		if height < num {
			height = num
		}
	}

	result := 0

	for h := 1; h <= height; h++ {

		countWaterTrap := 0
		lastBorderLeft := 0

		for w := 1; w <= width; w++ {
			remainingBlock := nums[w-1] - h

			// Find trap
			if remainingBlock < 0 && w-1 > 0 {
				countWaterTrap++
				continue
			}

			// Find position border left
			if remainingBlock >= 0 && w-1 == 0 {
				lastBorderLeft = w
			}

			// Find position border right
			if remainingBlock >= 0 && w-1 > 1 && w > lastBorderLeft && lastBorderLeft > 0 {
				result += countWaterTrap
				countWaterTrap = 0
				lastBorderLeft = w
			}
		}
	}

	fmt.Println("Result:", result)

	return result
}
