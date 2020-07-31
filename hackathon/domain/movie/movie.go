package movie

import (
	"github.com/jupemara/ddd-guys/hackathon/domain/movie/score"
)

type Movie struct {
	id              Id
	title           Title
	publishdPeriods []PublishPeriod
	scores          score.Scores
	length          Length
	labels          []Label
}

// TODO: Length

func (m Movie) Title() string {
	return m.title.Value()
}

func (m Movie) Labels() []Label {
	return m.labels
}

func (m Movie) Length() Length {
	return m.length
}

func (m Movie) Scores() score.Scores {
	return m.scores
}
