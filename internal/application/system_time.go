package application

import "time"

type SystemTime struct {
	time *time.Time
}

// NewSystemTime initializes new system time with current year, month, day, time 00:00
func NewSystemTime() *SystemTime {
	now := time.Now().UTC()
	newDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

	return &SystemTime{
		time: &newDate,
	}
}

// Time returns value of system time
func (t *SystemTime) Time() *time.Time {
	return t.time
}

func (t *SystemTime) Add(hours int) {
	tt := t.time.Add(time.Hour * time.Duration(hours))
	t.time = &tt
}
