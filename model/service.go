package model

import "errors"

type TodoService struct {
	repository TodoRepository
}

func NewTodoService(r TodoRepository, i RepositoryInitializer) *TodoService {
	service := &TodoService{repository: r}
	i.Initialize(r)
	return service
}

func (s *TodoService) FindAll() []*Todo {
	return s.repository.FindAll()
}

func (s *TodoService) FindById(id string) (*Todo, error) {
	return s.repository.FindById(id)
}

func (s *TodoService) CreateTodo(entry TodoNew) (*Todo, error) {

	t := &Todo{
		Item:      entry.Item,
		Completed: false,
	}

	return s.Save(t)
}

func (s *TodoService) UpdateTodo(id string, update *TodoUpdate) (bool, *Todo, error) {

	storedTodo, err := s.FindById(id)
	if err != nil {
		return false, nil, err
	}

	if storedTodo == nil {
		return false, nil, errors.New("notfound")
	}

	if update.Item != "" {
		// we only allow non-empty text change
		storedTodo.Item = update.Item
	}

	if update.MaybeCompleted.Valid {
		// only change completed status if we actually received and update for it
		storedTodo.Completed = update.MaybeCompleted.Bool
	}

	saved, err := s.Save(storedTodo)
	return true, saved, err
}

func (s *TodoService) Save(todo *Todo) (*Todo, error) {
	return s.repository.Save(todo)
}
