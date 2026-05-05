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

	lastStart := -1
	for _, in := range intervals {
		if lastStart == -1 {
			lastStart = in.Start
			continue
		}

		if lastStart >= in.Start {
			return false
		}

		lastStart = in.Start
	}

	lastEnd := -1
	for _, in := range intervals {
		if lastEnd == -1 {
			lastEnd = in.End
			continue
		}

		if lastEnd >= in.End || lastEnd > in.Start {
			return false
		}

		lastEnd = in.End
	}

	return true
}
