package model

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

func (s *TodoService) UpdateTodo(todo *Todo, update *TodoUpdate) (*Todo, error) {

	storedTodo, err := s.FindById(todo.ID)
	if err != nil {
		return nil, err
	}

	if update.Item != "" {
		// we only allow non-empty text change
		storedTodo.Item = update.Item
	}

	if update.MaybeCompleted.Valid {
		storedTodo.Completed = update.MaybeCompleted.Bool
	}

	return s.Save(storedTodo)
}

func (s *TodoService) Save(todo *Todo) (*Todo, error) {
	return s.repository.Save(todo)
}
