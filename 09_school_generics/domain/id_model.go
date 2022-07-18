package domain

type IDModeler interface {
	HasId() bool
	SetId(id int64)
}

type IDModel struct {
	ID *int64
}

func (im IDModel) HasId() bool {
	return im.ID != nil
}

func (im *IDModel) SetId(id int64) {
	im.ID = &id
}
