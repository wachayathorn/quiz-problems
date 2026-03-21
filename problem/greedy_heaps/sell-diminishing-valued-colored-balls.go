package greedy_heaps

import (
	"slices"
)

func MaxProfitV2(inventory []int, orders int) int {
	const MOD = 1_000_000_007
	slices.SortFunc(inventory, func(a, b int) int { return b - a })
	inventory = append(inventory, 0)

	profit := int64(0)
	n := len(inventory)

	for i := 0; i < n-1; i++ {
		if inventory[i] > inventory[i+1] {
			count := int64(i + 1)
			diff := int64(inventory[i] - inventory[i+1])
			canSell := count * diff

			if int64(orders) >= canSell {
				first := int64(inventory[i])
				last := int64(inventory[i+1] + 1)
				numRows := diff

				currentSum := (first + last) * numRows / 2
				profit = (profit + currentSum*count) % MOD
				orders -= int(canSell)
			} else {
				numFullRows := int64(orders) / count
				rem := int64(orders) % count

				first := int64(inventory[i])
				last := first - numFullRows + 1

				currentSum := (first + last) * numFullRows / 2
				profit = (profit + currentSum*count) % MOD

				profit = (profit + (last-1)*rem) % MOD
				orders = 0
				break
			}
		}
	}
	return int(profit)
}

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
