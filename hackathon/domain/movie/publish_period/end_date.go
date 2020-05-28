package publish_period

import (
	"errors"
	"time"
)

type EndDate struct {
	value time.Time
}

func (ed EndDate) Value() time.Time {
	return ed.value
}

func newEndDate(value string) (*EndDate, error) {
	timeValue, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return nil, errors.New("assertion error")
	}
	return &EndDate{value: timeValue}, nil
}
