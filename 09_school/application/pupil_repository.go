package application

import "school/domain"

type PupilRepository interface {
	Add(domain.Pupil) domain.Pupil
	All() domain.Pupils
	Find(id int64) *domain.Pupil
}
