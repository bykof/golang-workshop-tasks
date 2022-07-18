package infrastructure

import (
	"school/application"
	"school/domain"
)

type RealTeacherRepository struct {
	RealBaseRepository[*domain.Teacher]
}

var _ application.TeacherRepository = &RealTeacherRepository{}

func NewRealTeacherRepository() *RealTeacherRepository {
	return &RealTeacherRepository{
		RealBaseRepository: NewRealBaseRepository[*domain.Teacher](),
	}
}
