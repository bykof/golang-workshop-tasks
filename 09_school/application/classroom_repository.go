package application

import "school/domain"

type ClassRoomRepository interface {
	Add(domain.ClassRoom) domain.ClassRoom
	Find(id int64) *domain.ClassRoom
	All() domain.ClassRooms
}
