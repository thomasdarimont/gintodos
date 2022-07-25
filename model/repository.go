package model

type TodoRepository interface {
	FindAll() []*Todo

	FindById(id string) (*Todo, error)

	Save(entry *Todo) (*Todo, error)
}
