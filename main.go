package main

import (
	"fmt"

	"github.com/wachayathorn/quiz-problems/ploblem"
)

func main() {
	fmt.Println(ploblem.MinimumOperations([][]int{
		{3, 2, 1},
		{2, 1, 0},
		{1, 2, 3},
	}))
}
