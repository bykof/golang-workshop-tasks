package domain

import "fmt"

type Pupil struct {
	IDModel
	FirstName string
	LastName  string
	Grade     int64
}

type Pupils []Pupil

func (t Pupil) String() string {
	if !t.HasId() {
		return fmt.Sprintf(
			"Pupil %s %s",
			t.FirstName,
			t.LastName,
		)
	}
	return fmt.Sprintf(
		"Pupil #%d %s %s",
		*t.ID,
		t.FirstName,
		t.LastName,
	)
}
