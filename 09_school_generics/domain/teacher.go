package domain

import "fmt"

var _ IDModeler = &Teacher{}

type Teacher struct {
	IDModel
	FirstName    string
	LastName     string
	TeachingArea string
}

type Teachers []Teacher

func (t Teacher) String() string {
	if !t.HasId() {
		return fmt.Sprintf(
			"Teacher %s %s",
			t.FirstName,
			t.LastName,
		)
	}
	return fmt.Sprintf(
		"Teacher #%d %s %s",
		*t.ID,
		t.FirstName,
		t.LastName,
	)
}
