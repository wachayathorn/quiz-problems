package arrays_hashing

import "time"

type Booking struct {
	Start time.Time
	End   time.Time
}

func canBook(existing []Booking, newBooking Booking) bool {
	if !newBooking.End.After(newBooking.Start) {
		return false
	}
	for _, eb := range existing {
		if newBooking.Start.Before(eb.End) && newBooking.End.After(eb.Start) {
			return false
		}
	}
	return true
}
