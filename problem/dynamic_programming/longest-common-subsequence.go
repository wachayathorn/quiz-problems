package dynamic_programming

import (
	"math"
)

// LongestCommonSubsequence finds the length of the longest common subsequence between two strings.
// Time Complexity: O(m * n) - where m and n are the lengths of text1 and text2.
// Space Complexity: O(m * n) - for the DP table (can be optimized to O(min(m, n))).
func LongestCommonSubsequence(text1 string, text2 string) int {
	m := len(text1)
	n := len(text2)

	// Create a 2D DP table. dp[i][j] represents the LCS length of text1[:i] and text2[:j].
	// We use (m+1) x (n+1) to handle the empty string cases (base cases).
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Iterate through both strings
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// If characters match, LCS increases by 1 from the LCS of prefixes before these characters.
			// Formula: dp[i][j] = dp[i-1][j-1] + 1
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				// If characters don't match, take the maximum LCS by either:
				// 1. Excluding text1's current character: dp[i-1][j]
				// 2. Excluding text2's current character: dp[i][j-1]
				// Formula: dp[i][j] = max(dp[i-1][j], dp[i][j-1])
				dp[i][j] = int(math.Max(float64(dp[i-1][j]), float64(dp[i][j-1])))
			}
		}
	}

	// The final result is in the bottom-right cell
	return dp[m][n]
}
