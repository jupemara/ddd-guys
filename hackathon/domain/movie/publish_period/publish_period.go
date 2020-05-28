package publish_period

import "errors"

type PublishPeriod struct {
	contentsProvider ContentsProvider
	startDate        *StartDate
	endDate          *EndDate
}

func NewPublishPeriod(cp ContentsProvider, startDate string, endDate string) (*PublishPeriod, error) {
	for _, fn := range []publishPeriodConstructor{
		newPublishPeriodWithStartAndEndDate,
		newPublishPeriodWithStartDateOnly,
		newPublishPeriodWithEndDateOnly,
		newPublishPeriodWithoutDate,
	} {
		v, err := fn(cp, startDate, endDate)
		if err == nil {
			return v, nil
		}
	}
	return nil, errors.New("assertion error")
}

type publishPeriodConstructor func(cp ContentsProvider, startDate string, endDate string) (*PublishPeriod, error)

func newPublishPeriodWithStartAndEndDate(cp ContentsProvider, startDate string, endDate string) (*PublishPeriod, error) {
	start, err := newStartDate(startDate)
	if err != nil {
		return nil, err
	}
	end, err := newEndDate(endDate)
	if err != nil {
		return nil, err
	}
	if start.Value().After(end.Value()) {
		return nil, errors.New("assertion error")
	}
	return &PublishPeriod{contentsProvider: cp, startDate: start, endDate: end}, nil
}

func newPublishPeriodWithStartDateOnly(cp ContentsProvider, startDate string, endDate string) (*PublishPeriod, error) {
	start, err := newStartDate(startDate)
	if err != nil {
		return nil, err
	}
	return &PublishPeriod{contentsProvider: cp, startDate: start, endDate: nil}, nil
}

func newPublishPeriodWithEndDateOnly(cp ContentsProvider, startDate string, endDate string) (*PublishPeriod, error) {
	end, err := newEndDate(endDate)
	if err != nil {
		return nil, err
	}
	return &PublishPeriod{contentsProvider: cp, startDate: nil, endDate: end}, nil
}

func newPublishPeriodWithoutDate(cp ContentsProvider, startDate string, endDate string) (*PublishPeriod, error) {
	return &PublishPeriod{contentsProvider: cp, startDate: nil, endDate: nil}, nil
}
