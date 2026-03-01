package ploblem

import "strings"

// Input: word = "aabb"
// Output: 9
// Explanation: The nine wonderful substrings are underlined below:
// - "aabb" -> "a"
// - "aabb" -> "aa"
// - "aabb" -> "aab"
// - "aabb" -> "aabb"
// - "aabb" -> "a"
// - "aabb" -> "abb"
// - "aabb" -> "b"
// - "aabb" -> "bb"
// - "aabb" -> "b"
func WonderfulSubstrings(word string) int64 {
	result := 0
	chars := strings.Split(word, "")
	for _, c := range chars {
		alphabet := make(map[string]int)
		swicth, isExisting := alphabet[c]
		if !isExisting {
			alphabet[c] = 0
		}

		if swicth == 0 {
			alphabet[c] = 1
		} else {
			alphabet[c] = 0
		}

		sum := 0
		for _, a := range alphabet {
			sum += a
		}

		if sum <= 1 {
			result++
		}
	}

	return int64(result)
}
