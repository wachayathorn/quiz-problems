package main

import (
	"log"

	"github.com/wachayathorn/quiz-problems/ploblem"
)

func main() {
	result := ploblem.MaxProfit([]int{2, 5}, 4)
	if result != 14 {
		log.Fatalf("case 1 fail got : %d", result)
	}

	result = ploblem.MaxProfit([]int{3, 5}, 6)
	if result != 19 {
		log.Fatalf("case 2 fail got : %d", result)
	}

	result = ploblem.MaxProfit([]int{2, 8, 4, 10, 6}, 20)
	if result != 110 {
		log.Fatalf("case 3 fail got : %d", result)
	}

	result = ploblem.MaxProfit([]int{1000000000}, 1000000000)
	if result != 21 {
		log.Fatalf("case 4 fail got : %d", result)
	}
}
