package main

import (
	"fmt"
	"log"

	"github.com/wachayathorn/quiz-problems/ploblem"
)

func main() {
	result := ploblem.MinDeletions("aab")
	if result != 0 {
		log.Fatalf("case 1 error got : %d", result)
	} else {
		fmt.Println("case 1 success")
	}

	result = ploblem.MinDeletions("aaabbbcc")
	if result != 2 {
		log.Fatalf("case 2 error got : %d", result)
	} else {
		fmt.Println("case 2 success")
	}

	result = ploblem.MinDeletions("ceabaacb")
	if result != 2 {
		log.Fatalf("case 3 error got : %d", result)
	} else {
		fmt.Println("case 3 success")
	}

	result = ploblem.MinDeletions("accdcdadddbaadbc")
	if result != 1 {
		log.Fatalf("case 4 error got : %d", result)
	} else {
		fmt.Println("case 4 success")
	}

	result = ploblem.MinDeletions("abcdefghijklmnopqrstuvwxwzabcdefghijklmnopqrstuvwxwzabcdefghijklmnopqrstuvwxwzabcdefghijklmnopqrstuvwxwzabcdefghijklmnopqrstuvwxwzabcdefghijklmnopqrstuvwxwzabcdefghijklmnopqrstuvwxwzabcdefghijklmnopqrstuvwxwzabcdefghijklmnopqrstuvwxwzabcdefghijklmnopqrstuvwxwzabcdefghijklmnopqrstuvwxwzabcdefghijklmnopqrstuvwxwzabcdefghijklmnopqrstuvwxwzabcdefghijklmnopqrstuvwxwzabcdefghijklmnopqrstuvwxwzabcdefghijklmnopqrstuvwxwzabcdefghijklmnopqrstuvwxwzabcdefghijklmnopqrstuvwxwzabcdefghijklmnopqrstuvwxwzabcdefghijklmnopqrstuvwxwzabcdefghijklmnopqrstuvwxwzabcdefghijklmnopqrstuvwxwzabcdefghijklmnopqrstuvwxwzabcdefghijklmnopqrstuvwxwzabcdefghijklmnopqrstuvwxwzabcdefghijklmnopqrstuvwxwz")
	if result != 276 {
		log.Fatalf("case 5 error got : %d", result)
	} else {
		fmt.Println("case 5 success")
	}

	result = ploblem.MinDeletions("bbcebab")
	if result != 2 {
		log.Fatalf("case 6 error got : %d", result)
	} else {
		fmt.Println("case 6 success")
	}

}
