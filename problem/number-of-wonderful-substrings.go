package problem

func WonderfulSubstrings(word string) int64 {
	var result int64 = 0
	n := len(word)

	for i := 0; i < n; i++ {
		alphabet := make([]int, 10)
		oddCount := 0

		for j := i; j < n; j++ {
			charIdx := word[j] - 'a'

			if alphabet[charIdx] == 0 {
				alphabet[charIdx] = 1
				oddCount++ // ไฟดวงนี้สว่างขึ้นมา
			} else {
				alphabet[charIdx] = 0
				oddCount-- // ไฟดวงนี้ดับลง
			}

			if oddCount <= 1 {
				result++
			}
		}
	}

	return result
}
