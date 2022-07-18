package infrastructure

import (
	"school/application"
	"school/domain"
)

type RealPupilRepository struct {
	RealBaseRepository[*domain.Pupil]
}

var _ application.PupilRepository = &RealPupilRepository{}

func NewRealPupilRepository() *RealPupilRepository {
	return &RealPupilRepository{
		RealBaseRepository: NewRealBaseRepository[*domain.Pupil](),
	}
}
