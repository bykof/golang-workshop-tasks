package infrastructure

import (
	"errors"
	"school/application"
	"school/domain"
)

type RealLessonRepository struct {
	RealBaseRepository[*domain.Lesson]
}

var _ application.LessonRepository = &RealLessonRepository{}

func NewRealLessonRepository() *RealLessonRepository {
	return &RealLessonRepository{
		RealBaseRepository: NewRealBaseRepository[*domain.Lesson](),
	}
}

func (rcrr *RealLessonRepository) Update(lesson domain.Lesson) error {
	if !lesson.HasId() {
		return errors.New("lesson has no ID")
	}
	rcrr.storage[*lesson.ID] = &lesson
	return nil
}

func (rcrr *RealLessonRepository) FilterByClassRoom(classRoomId int64) domain.Lessons {
	var lessons domain.Lessons

	for _, lesson := range rcrr.storage {
		if lesson.HasId() && *lesson.ClassRoom.ID == classRoomId {
			lessons = append(lessons, *lesson)
		}
	}
	return lessons
}
