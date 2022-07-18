package application

import "school/domain"

type ClassRoomRepository interface {
	BaseRepository[*domain.ClassRoom]
}
