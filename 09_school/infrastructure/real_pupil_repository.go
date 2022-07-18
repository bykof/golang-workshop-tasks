package infrastructure

import (
	"school/application"
	"school/domain"
)

type RealPupilRepository struct {
	idCounter int64
	storage   map[int64]domain.Pupil
}

var _ application.PupilRepository = &RealPupilRepository{}

func NewRealPupilRepository() *RealPupilRepository {
	return &RealPupilRepository{
		idCounter: 1,
		storage:   make(map[int64]domain.Pupil),
	}
}

func (rcrr *RealPupilRepository) nextFreeId() int64 {
	value := rcrr.idCounter
	rcrr.idCounter++
	return value
}

func (rcrr *RealPupilRepository) Add(pupil domain.Pupil) domain.Pupil {
	nextFreeId := rcrr.nextFreeId()
	pupil.ID = &nextFreeId
	rcrr.storage[nextFreeId] = pupil
	return pupil
}

func (rcrr *RealPupilRepository) All() domain.Pupils {
	var pupils domain.Pupils
	for _, pupil := range rcrr.storage {
		pupils = append(pupils, pupil)
	}
	return pupils
}

func (rcrr *RealPupilRepository) Find(id int64) *domain.Pupil {
	pupil, ok := rcrr.storage[id]
	if !ok {
		return nil
	}
	return &pupil
}
