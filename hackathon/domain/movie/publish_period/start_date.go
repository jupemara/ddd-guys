package publish_period

import (
	"errors"
	"time"
)

type StartDate struct {
	value time.Time
}

func (sd StartDate) Value() time.Time {
	return sd.value
}

func newStartDate(value string) (*StartDate, error) {
	timeValue, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return nil, errors.New("assertion error")
	}
	return &StartDate{value: timeValue}, nil
}
