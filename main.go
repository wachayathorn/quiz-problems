package main

import (
	"fmt"

	"github.com/wachayathorn/quiz-problems/problem/arrays_hashing"
)

func main() {
	// [(0,30),(5,10),(15,20)]
	fmt.Println(arrays_hashing.CanAttendMeetings([]arrays_hashing.Interval{
		{Start: 0, End: 30},
		{Start: 5, End: 10},
		{Start: 15, End: 20},
	}))

	// [(5,8),(9,15)]
	fmt.Println(arrays_hashing.CanAttendMeetings([]arrays_hashing.Interval{
		{Start: 5, End: 8},
		{Start: 9, End: 15},
	}))

	// [(1,10),(9,20),(19,30),(29,40),(39,50)]
	fmt.Println(arrays_hashing.CanAttendMeetings([]arrays_hashing.Interval{
		{Start: 5, End: 8},
		{Start: 9, End: 15},
	}))
}
