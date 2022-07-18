package infrastructure

import (
	"school/application"
	"school/domain"
)

type RealTeacherRepository struct {
	idCounter int64
	storage   map[int64]domain.Teacher
}

var _ application.TeacherRepository = &RealTeacherRepository{}

func NewRealTeacherRepository() *RealTeacherRepository {
	return &RealTeacherRepository{
		idCounter: 1,
		storage:   make(map[int64]domain.Teacher),
	}
}

func (rcrr *RealTeacherRepository) nextFreeId() int64 {
	value := rcrr.idCounter
	rcrr.idCounter++
	return value
}

func (rcrr *RealTeacherRepository) Find(id int64) *domain.Teacher {
	teacher, ok := rcrr.storage[id]
	if !ok {
		return nil
	}
	return &teacher
}

func (rcrr *RealTeacherRepository) Add(teacher domain.Teacher) domain.Teacher {
	nextFreeId := rcrr.nextFreeId()
	teacher.ID = &nextFreeId
	rcrr.storage[nextFreeId] = teacher
	return teacher
}

func (rcrr *RealTeacherRepository) All() domain.Teachers {
	var teachers domain.Teachers
	for _, teacher := range rcrr.storage {
		teachers = append(teachers, teacher)
	}
	return teachers
}
