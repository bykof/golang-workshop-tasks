package application

import "school/domain"

type BaseRepository[T domain.IDModeler] interface {
	Add(item T) T
	Find(id int64) *T
	All() []T
}
