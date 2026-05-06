package arrays_hashing

import "slices"

func groupAnagrams(strs []string) [][]string {
	charsWithStrs := map[string][]string{}
	for _, s := range strs {
		runes := []rune(s)
		slices.Sort(runes)
		sortedStr := string(runes)
		charsWithStrs[sortedStr] = append(charsWithStrs[sortedStr], s)
	}
	results := [][]string{}
	for _, strs := range charsWithStrs {
		results = append(results, strs)
	}
	return results
}
