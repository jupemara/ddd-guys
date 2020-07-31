package score

import "fmt"

const yahooMaxScore = 5.0

type YahooScore struct {
	scoreProvider ScoreProvider
	value         float64
	maxValue      float64
	turnout       int
}

func NewYahooScore(scoreProvider int, value float64, turnout int) (Score, error) {

	for _, v := range []bool{
		scoreProvider == YAHOO,
		value > 0,
		value <= yahooMaxScore,
		turnout > 0,
	} {
		if !v {
			// TODO(makocchi-git) refactor error msg
			return nil, fmt.Errorf("assertion error")
		}
	}
	return &YahooScore{
		scoreProvider: scoreProvider,
		value:         value,
		maxValue:      yahooMaxScore,
		turnout:       turnout,
	}, nil
}

func (y YahooScore) ScoreProvider() ScoreProvider {
	return y.scoreProvider
}

func (y YahooScore) Value() float64 {
	return y.value
}

func (y YahooScore) MaxValue() float64 {
	return y.maxValue
}

func (y YahooScore) Turnout() int {
	return y.turnout
}
