package increasetime

import "fmt"

type response struct {
	Time string
}

func NewResponse(time string) *response {
	return &response{Time: time}
}

func (r *response) String() string {
	return fmt.Sprintf("Time is %s", r.Time)
}
