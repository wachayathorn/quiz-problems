package ploblem

// Input: grid = [[3,2],[1,3],[3,4],[0,1]]
// Output: 15
// Explanation:
// To make the 0th column strictly increasing, we can apply 3 operations on grid[1][0], 2 operations on grid[2][0], and 6 operations on grid[3][0].
// To make the 1st column strictly increasing, we can apply 4 operations on grid[3][1].
func MinimumOperations(grid [][]int) int {
	column := len(grid[0])
	result := 0
	for c := 0; c < column; c++ {
		for g := range grid {
			// skip first
			if g == 0 {
				continue
			}
			// check current less than previous
			for grid[g][c] <= grid[g-1][c] {
				grid[g][c]++
				result++
			}
		}
	}
	return result
}
