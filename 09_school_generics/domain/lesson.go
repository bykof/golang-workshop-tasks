package domain

import "fmt"

var _ IDModeler = &Lesson{}

type Lesson struct {
	IDModel
	Name      string
	Semester  string
	ClassRoom ClassRoom
	Teacher   Teacher
	Pupils    Pupils
}

type Lessons []Lesson

func (t Lesson) String() string {
	if !t.HasId() {
		return fmt.Sprintf(
			"Lesson %s %s",
			t.Name,
			t.Semester,
		)
	}
	return fmt.Sprintf(
		"Lesson #%d %s %s",
		*t.ID,
		t.Name,
		t.Semester,
	)
}
