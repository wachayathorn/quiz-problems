package strings

import "strings"

// Replace All ?'s to Avoid Consecutive Repeating Characters.
// ?zs -> azs
func ModifyString(str string) string {
	var a2zList = []string{
		"a",
		"b",
		"c",
		"d",
		"e",
		"f",
		"g",
		"h",
		"i",
		"j",
		"k",
		"l",
		"m",
		"n",
		"o",
		"p",
		"q",
		"r",
		"s",
		"t",
		"u",
		"v",
		"w",
		"x",
		"y",
		"z",
	}
	strs := strings.Split(str, "")

	if len(strs) == 1 && strs[0] == "?" {
		return "a"
	}

	for i, s := range strs {

		// Skip is not "?"
		if s != "?" {
			continue
		}

		// get prev string
		prevString := ""
		if i > 0 {
			prevString = strs[i-1]
			for _, c := range a2zList {
				if prevString != c {
					strs[i] = c
					break
				}
			}
		}

		// get next string
		nextString := ""
		if i < len(strs)-1 {
			nextString = strs[i+1]
			for _, c := range a2zList {
				if nextString != c && c != prevString {
					strs[i] = c
					break
				}
			}
		}
	}

	return strings.Join(strs, "")
}
