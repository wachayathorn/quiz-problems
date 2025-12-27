package ploblem

import (
	"fmt"
)

func ArrayRotations(arr []int, k int) {
	for i := 1; i <= k; i++ {
		for j := len(arr) - 1; j > 0; j-- {
			temp := arr[j]
			arr[j] = arr[j-1]
			arr[j-1] = temp
		}
	}
	fmt.Println("Result : ", arr)
}
