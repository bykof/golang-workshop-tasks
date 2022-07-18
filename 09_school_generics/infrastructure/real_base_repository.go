package infrastructure

import (
	"school/application"
	"school/domain"
)

type RealBaseRepository[T domain.IDModeler] struct {
	idCounter int64
	storage   map[int64]T
}

var _ application.BaseRepository[domain.IDModeler] = &RealBaseRepository[domain.IDModeler]{}

func (rbr *RealBaseRepository[T]) nextFreeId() int64 {
	value := rbr.idCounter
	rbr.idCounter++
	return value
}

func NewRealBaseRepository[T domain.IDModeler]() RealBaseRepository[T] {
	return RealBaseRepository[T]{
		idCounter: 1,
		storage:   make(map[int64]T),
	}
}

func (rbr *RealBaseRepository[T]) Add(item T) T {
	nextFreeId := rbr.nextFreeId()
	item.SetId(nextFreeId)
	rbr.storage[nextFreeId] = item
	return item
}

func (rbr *RealBaseRepository[T]) All() []T {
	var items []T
	for _, item := range rbr.storage {
		items = append(items, item)
	}
	return items
}

func (rbr *RealBaseRepository[T]) Find(id int64) *T {
	item, ok := rbr.storage[id]
	if !ok {
		return nil
	}
	return &item
}
