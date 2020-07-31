package publish_period

import "errors"

type ContentsProvider struct {
	value string
}

func (cp ContentsProvider) Value() string {
	return cp.value
}

func NewContentsProvider(value string) (*ContentsProvider, error) {
	for _, unsatisfied := range []bool{
		len(value) < 2,
		len(value) > 16,
	} {
		if unsatisfied {
			return nil, errors.New("doesn't meet the criteria.")
		}
	}
	return &ContentsProvider{value: value}, nil
}
