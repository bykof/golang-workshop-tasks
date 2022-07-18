package domain

import "fmt"

var _ IDModeler = &ClassRoom{}

type ClassRoom struct {
	IDModel
	Number   string
	Location string
}

type ClassRooms []ClassRoom

func (t ClassRoom) String() string {
	if !t.HasId() {
		return fmt.Sprintf(
			"Classroom %s %s",
			t.Number,
			t.Location,
		)
	}
	return fmt.Sprintf(
		"Classroom #%d %s %s",
		*t.ID,
		t.Number,
		t.Location,
	)
}
