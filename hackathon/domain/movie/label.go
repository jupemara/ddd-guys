package movie

import "errors"

const (
	MinLabelLength = 2
	MaxLabelLength = 16
)

func (m Movie) Labels() []Label {
	return m.labels
}

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
			return nil, errors.New(
				"label must be between %d and %d",
				MinLabelLength,
				MaxLabelLength,
			)
		}
	}
	return &Label{value: value}, nil
}
