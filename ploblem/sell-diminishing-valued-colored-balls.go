package ploblem

import (
	"slices"
)

// [4,5,6,4,2] , o = 20
// expect = 110
func MaxProfit(inventories []int, orders int) int {
	const MOD = 1_000_000_007

	slices.SortFunc(inventories, func(a, b int) int {
		return b - a
	})
	profit := 0
	for i := 0; i < orders; i++ {
		for currentPosition, currentBall := range inventories {
			nextPosition := currentPosition + 1
			nextBall := 0
			if nextPosition < len(inventories) {
				nextBall = inventories[nextPosition]
			}
			if currentBall >= nextBall {
				inventories[currentPosition]--
				profit += currentBall
				if nextPosition < len(inventories) && inventories[currentPosition] < inventories[nextPosition] {
					slices.SortFunc(inventories, func(a, b int) int {
						return b - a
					})
				}
				break
			}
		}
	}
	return int(profit % MOD)
}

// func MaxProfit0.1(inventories []int, orders int) int {
// 	ballsValue := []int{}
// 	for _, inventory := range inventories {
// 		for i := 1; i <= inventory; i++ {
// 			ballsValue = append(ballsValue, i)
// 		}
// 	}

// 	slices.SortFunc(ballsValue, func(a, b int) int {
// 		return b - a
// 	})

// 	profit := 0
// 	for i := 0; i < orders; i++ {
// 		profit += ballsValue[i]
// 	}

// 	return profit
// }
