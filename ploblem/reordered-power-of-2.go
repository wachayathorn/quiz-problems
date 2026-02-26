package ploblem

import (
	"sort"
	"strconv"
	"strings"
)

func ReorderedPowerOf2(n int) bool {
	nDigits := sortDigits(n)

	for i := 0; i < 31; i++ {
		powerOf2 := 1 << uint(i) // 00001 << i shifts left to get 2^i

		if sortDigits(powerOf2) == nDigits {
			return true
		}
	}

	return false
}

func sortDigits(num int) string {
	str := strconv.Itoa(num)
	digits := strings.Split(str, "")
	sort.Strings(digits)
	return strings.Join(digits, "")
}
