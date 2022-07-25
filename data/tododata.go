package data

import "gintodos/model"

type todoExampleData struct{}

var _ model.RepositoryInitializer = (*todoExampleData)(nil)

func NewExampleDataInitializer() model.RepositoryInitializer {
	return &todoExampleData{}
}

// Initializes the given TodoRepository with initial ToDo items
func (t *todoExampleData) Initialize(r model.TodoRepository) {

	for _, newTodo := range []*model.Todo{
		{Item: "Item 1", Completed: false},
		{Item: "Item 2", Completed: false},
		{Item: "Item 3", Completed: false},
	} {
		_, err := r.Save(newTodo)
		if err != nil {
			// ignore error if todo could not be stored
			continue
		}
	}
}
