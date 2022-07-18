package application

import "school/domain"

type PupilRepository interface {
	BaseRepository[*domain.Pupil]
}
