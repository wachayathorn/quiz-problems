package main

import (
	"log"

	"github.com/wachayathorn/quiz-problems/problem/dynamic_programming"
)

func main() {
	result := dynamic_programming.LongestCommonSubsequence("cat", "crabt")
	if result != 3 {
		log.Fatalf("case 1 error got : %d", result)
	}

	result = dynamic_programming.LongestCommonSubsequence("abcd", "abcd")
	if result != 4 {
		log.Fatalf("case 2 error got : %d", result)
	}

	result = dynamic_programming.LongestCommonSubsequence("abcd", "efgh")
	if result != 0 {
		log.Fatalf("case 3 error got : %d", result)
	}

	result = dynamic_programming.LongestCommonSubsequence("crabt", "car")
	if result != 2 {
		log.Fatalf("case 4 error got : %d", result)
	}

	result = dynamic_programming.LongestCommonSubsequence("bsbininm", "jmjkbkjkv")
	if result != 1 {
		log.Fatalf("case 4 error got : %d", result)
	}
}
