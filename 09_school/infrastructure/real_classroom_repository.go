package infrastructure

import (
	"school/application"
	"school/domain"
)

type RealClassRoomRepository struct {
	idCounter int64
	storage   map[int64]domain.ClassRoom
}

var _ application.ClassRoomRepository = &RealClassRoomRepository{}

func NewRealClassRoomRepository() *RealClassRoomRepository {
	return &RealClassRoomRepository{
		idCounter: 1,
		storage:   make(map[int64]domain.ClassRoom),
	}
}

func (rcrr *RealClassRoomRepository) nextFreeId() int64 {
	value := rcrr.idCounter
	rcrr.idCounter++
	return value
}

func (rcrr *RealClassRoomRepository) Find(id int64) *domain.ClassRoom {
	classRoom, ok := rcrr.storage[id]
	if !ok {
		return nil
	}
	return &classRoom
}

func (rcrr *RealClassRoomRepository) Add(classRoom domain.ClassRoom) domain.ClassRoom {
	nextFreeId := rcrr.nextFreeId()
	classRoom.ID = &nextFreeId
	rcrr.storage[nextFreeId] = classRoom
	return classRoom
}

func (rcrr *RealClassRoomRepository) All() domain.ClassRooms {
	var classRooms domain.ClassRooms
	for _, classRoom := range rcrr.storage {
		classRooms = append(classRooms, classRoom)
	}
	return classRooms
}
