package store

import (
	"errors"
	"gintodos/model"
	"strconv"
	"time"
)

type memoryStore struct {
	entries []*model.Todo
}

// ensure memoryStore implements TodoRepository interface
var _ model.TodoRepository = (*memoryStore)(nil)

func NewMemoryTodoStore() *memoryStore {
	return &memoryStore{
		entries: []*model.Todo{},
	}
}

// FindAll returns a slice of ToDos
func (s *memoryStore) FindAll() []*model.Todo {
	return s.entries
}

// FindById returns a ToDo item found by ID or nil if missing
func (s *memoryStore) FindById(id string) (*model.Todo, error) {

	for _, t := range s.entries {
		if t.ID == id {
			return t, nil
		}
	}

	return nil, errors.New("not found")
}

// Save saves a new ToDo or updates an existing ToDo
func (s *memoryStore) Save(t *model.Todo) (*model.Todo, error) {

	now := time.Now()
	if t.ID == "" {
		// insert new todo
		t.ID = strconv.FormatInt(now.Unix(), 10)
		t.CreatedAt = now
		t.ModifiedAt = now
		s.entries = append(s.entries, t)
		return t, nil
	}

	// try to update existing todo
	todo, err := s.FindById(t.ID)
	if err != nil {
		// could not find existing todo
		return nil, errors.New("notfound")
	}

	// found update existing todo
	todo.Item = t.Item
	todo.Completed = t.Completed
	todo.ModifiedAt = now
	return todo, nil
}
