package arrays_hashing

import (
	"cmp"
	"slices"
)

type Interval struct {
	Start int
	End   int
}

func CanAttendMeetings(intervals []Interval) bool {
	slices.SortFunc(intervals, func(a, b Interval) int {
		return cmp.Compare(a.Start, b.Start)
	})
	for i := 1; i < len(intervals); i++ {
		if intervals[i].Start < intervals[i-1].End {
			return false
		}
	}
	return true
}
