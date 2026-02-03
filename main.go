package main

import (
	"fmt"

	"github.com/wachayathorn/quiz-problems/ploblem"
)

func main() {
	num := 246
	result := ploblem.ReorderedPowerOf2(num)
	fmt.Printf("Can %d be reordered to form a power of 2? %v\n", num, result)
}
