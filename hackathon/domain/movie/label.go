package movie

import "fmt"

const (
	MinLabelLength = 2
	MaxLabelLength = 16
)

type Label struct {
	value string
}

func (l Label) Value() string {
	return l.value
}

func NewLabel(value string) (*Label, error) {
	// label must be between MinLabelLength and MaxLabelLength
	for _, unsatisfied := range []bool{
		len(value) < MinLabelLength,
		len(value) > MaxLabelLength,
	} {
		if unsatisfied {
			return nil, fmt.Errorf(
				"label must be between %d and %d",
				MinLabelLength,
				MaxLabelLength,
			)
		}
	}
	return &Label{value: value}, nil
}
