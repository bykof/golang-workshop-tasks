package application

import "school/domain"

type TeacherRepository interface {
	Add(domain.Teacher) domain.Teacher
	Find(id int64) *domain.Teacher
	All() domain.Teachers
}
