package ploblem

// A string s is called good if there are no two different characters in s that have the same frequency.
// Given a string s, return the minimum number of characters you need to delete to make s good.
// The frequency of a character in a string is the number of times it appears in the string. For example, in the string "aab", the frequency of 'a' is 2, while the frequency of 'b' is 1.
func MinDeletions(s string) int {
	uniqueCharWithValue := make(map[rune]int)
	for _, r := range s {
		uniqueCharWithValue[r]++
	}

	deleted := 0
	uniqueValue := make(map[int]bool)

	for _, value := range uniqueCharWithValue {
		freq := value
		for freq > 0 {
			_, isExisting := uniqueValue[freq]
			if !isExisting {
				uniqueValue[freq] = true
				break
			}
			freq--
			deleted++
		}
	}

	return deleted

	// strWithCount := make(map[string]int)
	// for _, c := range s {
	// 	_, ok := strWithCount[string(c)]
	// 	if !ok {
	// 		strWithCount[string(c)] = 1
	// 	} else {
	// 		strWithCount[string(c)]++
	// 	}
	// }

	// type KeyValuePair struct {
	// 	Key   string
	// 	Value int
	// }

	// var pairs []KeyValuePair
	// for k, v := range strWithCount {
	// 	pairs = append(pairs, KeyValuePair{k, v})
	// }

	// sort.Slice(pairs, func(i, j int) bool {
	// 	return pairs[i].Value < pairs[j].Value
	// })

	// uniqueValue := make(map[int]bool)

	// deleted := 0
	// for i := 0; i < len(pairs); i++ {
	// 	isGo := true
	// 	for isGo {
	// 		_, ok := uniqueValue[pairs[i].Value]
	// 		if !ok {
	// 			uniqueValue[pairs[i].Value] = true
	// 			isGo = false
	// 		} else if pairs[i].Value == 0 {
	// 			isGo = false
	// 		} else {
	// 			pairs[i].Value--
	// 			deleted++
	// 		}
	// 	}
	// }

	// return deleted
}
