package infrastructure

import (
	"errors"
	"school/application"
	"school/domain"
)

type RealLessonRepository struct {
	idCounter int64
	storage   map[int64]domain.Lesson
}

var _ application.LessonRepository = &RealLessonRepository{}

func NewRealLessonRepository() *RealLessonRepository {
	return &RealLessonRepository{
		idCounter: 1,
		storage:   make(map[int64]domain.Lesson),
	}
}

func (rcrr *RealLessonRepository) nextFreeId() int64 {
	value := rcrr.idCounter
	rcrr.idCounter++
	return value
}

func (rcrr *RealLessonRepository) Add(lesson domain.Lesson) domain.Lesson {
	nextFreeId := rcrr.nextFreeId()
	lesson.ID = &nextFreeId
	rcrr.storage[nextFreeId] = lesson
	return lesson
}

func (rcrr *RealLessonRepository) All() domain.Lessons {
	var lessons domain.Lessons
	for _, lesson := range rcrr.storage {
		lessons = append(lessons, lesson)
	}
	return lessons
}

func (rcrr *RealLessonRepository) Update(lesson domain.Lesson) error {
	if !lesson.HasId() {
		return errors.New("lesson has no ID")
	}
	rcrr.storage[*lesson.ID] = lesson
	return nil
}

func (rcrr *RealLessonRepository) Find(id int64) *domain.Lesson {
	lesson, ok := rcrr.storage[id]
	if !ok {
		return nil
	}
	return &lesson
}

func (rcrr *RealLessonRepository) FilterByClassRoom(classRoomId int64) domain.Lessons {
	var lessons domain.Lessons

	for _, lesson := range rcrr.storage {
		if lesson.HasId() && *lesson.ClassRoom.ID == classRoomId {
			lessons = append(lessons, lesson)
		}
	}
	return lessons
}
