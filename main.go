package main

import (
	"fmt"

	"github.com/wachayathorn/quiz-problems/problem/trees"
)

func main() {
	fmt.Println(trees.KthSmallest([]int{4, 3, 5, 2, 0}, 4))
	fmt.Println(trees.KthSmallest([]int{2, 1, 3}, 1))
}
