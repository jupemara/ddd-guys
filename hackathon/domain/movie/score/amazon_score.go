package score

import "fmt"

const amazonMaxScore = 5.0

type AmazonScore struct {
	scoreProvider ScoreProvider
	value         float64
	maxValue      float64
	turnout       int
}

func NewAmazonScore(scoreProvider int, value float64, turnout int) (Score, error) {

	for _, v := range []bool{
		scoreProvider == AMAZON,
		value > 0,
		value <= amazonMaxScore,
		turnout > 0,
	} {
		if !v {
			// TODO(makocchi-git) refactor error msg
			return nil, fmt.Errorf("assertion error")
		}
	}
	return &AmazonScore{
		scoreProvider: scoreProvider,
		value:         value,
		maxValue:      amazonMaxScore,
		turnout:       turnout,
	}, nil
}

func (a AmazonScore) ScoreProvider() ScoreProvider {
	return a.scoreProvider
}

func (a AmazonScore) Value() float64 {
	return a.value
}

func (a AmazonScore) MaxValue() float64 {
	return a.maxValue
}

func (a AmazonScore) Turnout() int {
	return a.turnout
}
