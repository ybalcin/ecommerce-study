package utility

import "time"

// DropMillisecond drops milliseconds of time and gets date, hours, minutes
func DropMillisecond(t time.Time) time.Time {
	newDate := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, time.UTC)

	return newDate
}
