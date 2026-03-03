package ploblem

// Given a string s, remove duplicate letters so that every letter appears once and only once. You must make sure your result is the smallest in lexicographical order among all possible results.
// Input: s = "cbacdcbc"
// Output: "acdb"
func RemoveDuplicateLetters(s string) string {
	result := ""

	// remove duplicated letter
	alphabet := make(map[rune]rune)
	for _, c := range s {
		charIdx := c - 'a'
		alphabet[c] = charIdx
	}

	bags := [][]rune{}

	return result
}
