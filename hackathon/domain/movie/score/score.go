package score

type IScore interface {
	ScoreProvider() ScoreProvider
	Value() float64
	MaxValue() float64
	Turnout() int
}
