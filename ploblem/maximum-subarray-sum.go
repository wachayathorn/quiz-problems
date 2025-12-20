package ploblem

import "math"

func MaxSubarraySum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	// 1. กำหนดค่าเริ่มต้น
	// global_max คือผลรวมสูงสุดที่เราเคยเจอมา
	// current_max คือผลรวมสูงสุดของ subarray ที่จบลงที่ตำแหน่งปัจจุบัน

	// เราใช้ Element แรกเป็นค่าเริ่มต้นของทั้งสองตัวแปร
	globalMax := arr[0]
	currentMax := arr[0]

	// เริ่มวนลูปตั้งแต่ Element ตัวที่ 2 (Index 1)
	for i := 1; i < len(arr); i++ {
		element := arr[i]

		// Logic 1: ตัดสินใจว่าจะเริ่มต้น Subarray ใหม่ หรือ ขยาย Subarray เดิม
		// ถ้า element ปัจจุบัน > (element ปัจจุบัน + currentMax ก่อนหน้า)
		// แสดงว่าเราควรเริ่มต้น Subarray ใหม่ด้วย element ตัวนี้
		currentMax = int(math.Max(float64(element), float64(element+currentMax)))

		// Logic 2: อัพเดทผลรวมสูงสุดโดยรวม (globalMax)
		// เปรียบเทียบผลรวมสูงสุดที่จบลง ณ ตำแหน่ง i กับผลรวมสูงสุดทั้งหมดที่เคยเจอ
		globalMax = int(math.Max(float64(globalMax), float64(currentMax)))
	}

	return globalMax
}
