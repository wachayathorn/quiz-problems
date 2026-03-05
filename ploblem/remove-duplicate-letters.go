package ploblem

import "strings"

func RemoveDuplicateLetters(s string) string {
	seen := make(map[byte]int)
	for _, char := range s {
		charIdx := byte(char)
		_, ok := seen[charIdx]
		if !ok {
			seen[charIdx] = 1
			continue
		}
		seen[charIdx]++
	}

	result := []byte{}
	for _, currentChar := range s {
		currentCharIdx := byte(currentChar)
		seen[currentCharIdx]--

		if len(result) == 0 {
			result = append(result, currentCharIdx)
			continue
		}

		if strings.Contains(string(result), string(currentChar)) {
			continue
		}

		latestCharIdx := result[len(result)-1]

		if currentCharIdx < latestCharIdx {
			temp := make([]byte, len(result))
			copy(temp, result)
			for i := len(temp) - 1; i >= 0; i-- {
				tempIdx := byte(temp[i])
				if seen[tempIdx] > 0 && currentCharIdx < tempIdx {
					result = result[:len(result)-1]
				} else {
					break
				}
			}
			result = append(result, currentCharIdx)
			continue
		}

		if currentCharIdx > latestCharIdx {
			if !strings.Contains(string(result), string(currentChar)) {
				result = append(result, currentCharIdx)
			}
			continue
		}
	}

	return string(result)
}
