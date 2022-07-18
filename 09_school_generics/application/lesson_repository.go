package application

import "school/domain"

type LessonRepository interface {
	BaseRepository[*domain.Lesson]
	Update(domain.Lesson) error
	FilterByClassRoom(classRoomId int64) domain.Lessons
}
