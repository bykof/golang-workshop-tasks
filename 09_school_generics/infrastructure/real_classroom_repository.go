package infrastructure

import (
	"school/application"
	"school/domain"
)

type RealClassRoomRepository struct {
	RealBaseRepository[*domain.ClassRoom]
}

var _ application.ClassRoomRepository = &RealClassRoomRepository{}

func NewRealClassRoomRepository() *RealClassRoomRepository {
	return &RealClassRoomRepository{
		RealBaseRepository: NewRealBaseRepository[*domain.ClassRoom](),
	}
}
