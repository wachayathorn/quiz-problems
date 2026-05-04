package strings

// "([{}])"
func IsValid(s string) bool {
	parenthesesMap := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}

	sMap := make(map[string]int, len(s))
	for _, r := range s {
		r := string(r)
		_, isExisting := sMap[r]
		if !isExisting {
			sMap[r] = 0
		} else {
			sMap[r]++
		}
	}

	for r, leftBucketTotal := range sMap {
		s := string(r)

		rightBacket, ok := parenthesesMap[s]
		if !ok {
			continue
		}

		rightBucketTotal, isExisting := sMap[string(rightBacket)]
		if !isExisting {
			return false
		}

		if leftBucketTotal != rightBucketTotal {
			return false
		}
	}

	return true
}

// You are given a string s consisting of the following characters: '(', ')', '{', '}', '[' and ']'.
// The input string s is valid if and only if:
// Every open bracket is closed by the same type of close bracket.
// Open brackets are closed in the correct order.
// Every close bracket has a corresponding open bracket of the same type.
// Return true if s is a valid string, and false otherwise.
func IsValidV2(s string) bool {
	leftBucket := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}
	rightBucket := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}
	stack := []string{}
	for _, r := range s {
		char := string(r)
		if _, ok := leftBucket[char]; ok {
			stack = append(stack, char)
		} else if _, ok := rightBucket[char]; ok {
			if len(stack) == 0 {
				return false
			}
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if last != rightBucket[char] {
				return false
			}
		}
	}
	return len(stack) == 0
}
