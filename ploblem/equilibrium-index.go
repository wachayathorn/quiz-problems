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

func EquilibriumIndexOptimized(nums []int) int {
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	leftSum := 0
	for i, num := range nums {
		// Right Sum คำนวณได้จาก: ผลรวมทั้งหมด - ผลรวมทางซ้าย - ตัวมันเอง
		rightSum := totalSum - leftSum - num

		if leftSum == rightSum {
			return i
		}

		// เตรียมค่า leftSum สำหรับ Index ถัดไป
		leftSum += num
	}

	return -1
}
