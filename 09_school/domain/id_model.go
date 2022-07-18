package domain

type IDModel struct {
	ID *int64
}

func (im IDModel) HasId() bool {
	return im.ID != nil
}
