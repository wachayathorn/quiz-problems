package ploblem

func FindMissingNumber(nums []int) int {
	newNums := nums
	for i := 0; i < len(nums); i++ {
		currentNum := i + 1
		for j, num := range newNums {
			if num == currentNum {
				newNums = append(newNums[:j], newNums[j+1:]...)
				break
			}
			if j == len(newNums)-1 {
				return currentNum
			}
		}
	}
	return 0
}

func MapFindMissingNumber(nums []int) int {
	numMap := make(map[int]bool)
	for _, num := range nums {
		numMap[num] = true
	}
	for i := 1; i <= len(nums)+1; i++ {
		if !numMap[i] {
			return i
		}
	}
	return 0
}

func FindMissingNumberBySum(nums []int) int {
	// 1. คำนวณค่า N (จำนวนตัวเลขทั้งหมดที่ควรจะมี)
	// ขนาดของ Array คือ n-1 ดังนั้น n = len(nums) + 1
	n := len(nums) + 1

	// 2. คำนวณผลรวมที่คาดหวัง (Expected Sum)
	// S_n = n * (n + 1) / 2
	// เราใช้ int64 เพื่อป้องกัน Integer Overflow หาก n มีขนาดใหญ่มาก
	expectedSum := (n * (n + 1)) / 2

	// 3. คำนวณผลรวมจริง (Actual Sum)
	// วนลูปบวก Element ทั้งหมดใน Array ที่มีอยู่
	actualSum := 0
	for _, num := range nums {
		actualSum += num
	}

	// 4. หาตัวเลขที่หายไป
	// Missing Number = Expected Sum - Actual Sum
	missingNumber := expectedSum - actualSum

	return missingNumber
}
