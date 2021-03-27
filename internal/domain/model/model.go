package model

type Model struct {
	id   int    // model id
	name string // model name
}

type ModelStorage interface {
	Create(b *Model) *Model
	Read(b *Model) *Model
	Update(b *Model) *Model
	Delete(id int) error
}
