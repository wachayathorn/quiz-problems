package problem

import (
	"fmt"
	"strings"
)

// successes
// The vowels are: 'u' (frequency 1), 'e' (frequency 2). The maximum frequency is 2.
// The consonants are: 's' (frequency 4), 'c' (frequency 2). The maximum frequency is 4.
// The output is 2 + 4 = 6.
func MaxFreqSum(s string) int {
	var (
		maxConsonantCount, maxVowelCount int
		vowels                           = map[string]int{
			"a": 0,
			"e": 0,
			"i": 0,
			"o": 0,
			"u": 0,
		}
		consonants = make(map[string]int)
	)

	strs := strings.Split(s, "")
	for _, s := range strs {
		_, isVowel := vowels[s]

		if isVowel {
			vowels[s]++
			if vowels[s] > maxVowelCount {
				maxVowelCount = vowels[s]
			}
			continue
		}

		_, haveConsonant := consonants[s]
		if haveConsonant {
			consonants[s]++
		} else {
			consonants[s] = 1
		}

		if consonants[s] > maxConsonantCount {
			maxConsonantCount = consonants[s]
		}
	}

	fmt.Println("maxConsonantCount is : ", maxConsonantCount)
	fmt.Println("maxVowelCount is : ", maxVowelCount)

	return maxConsonantCount + maxVowelCount
}
