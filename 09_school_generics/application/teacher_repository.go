package application

import "school/domain"

type TeacherRepository interface {
	BaseRepository[*domain.Teacher]
}
