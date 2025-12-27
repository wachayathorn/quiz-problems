package ploblem

import "fmt"

func ArrayRotations(arr []int, k int) {
	for i := 1; i <= k; i++ {
		lastValue := arr[len(arr)-1:][0]
		arr = arr[0 : len(arr)-1]
		arr = append([]int{lastValue}, arr...)
	}
	fmt.Println("Result : ", arr)
}
