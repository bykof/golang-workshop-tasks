package application

import "school/domain"

type LessonRepository interface {
	Add(domain.Lesson) domain.Lesson
	All() domain.Lessons
	Update(domain.Lesson) error
	Find(id int64) *domain.Lesson
	FilterByClassRoom(classRoomId int64) domain.Lessons
}
