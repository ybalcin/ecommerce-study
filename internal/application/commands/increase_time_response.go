package commands

import "fmt"

type increaseTimeResponse struct {
	Time string
}

func NewIncreaseTimeResponse(time string) *increaseTimeResponse {
	return &increaseTimeResponse{Time: time}
}

func (r *increaseTimeResponse) String() string {
	return fmt.Sprintf("Time is %s", r.Time)
}
