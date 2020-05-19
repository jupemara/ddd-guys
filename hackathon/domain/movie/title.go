package movie

import "fmt"

const (
	MinTitleLength = 1
	MaxTitleLength = 100
)

type Title struct {
	value string
}

func (t Title) Value() string {
	return t.value
}

func NewTitle(value string) (*Title, error) {
	// title must be between MinTitleLength and MaxTitleLength
	for _, unsatisfied := range []bool{
		len(value) < MinTitleLength,
		len(value) > MaxTitleLength,
	} {
		if unsatisfied {
			return nil, fmt.Errorf(
				"title must be between %d and %d",
				MinTitleLength,
				MaxTitleLength,
			)
		}
	}
	return &Title{value: value}, nil
}
