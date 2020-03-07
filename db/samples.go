package db

import (
	"github.com/ThomasHamilton2/todo-list/schema"
)

type Sample struct{}

func (s *Sample) Close() {}

func (s *Sample) Insert(todo *schema.Todo) (int, error) {
	return 0, nil
}

func (s *Sample) Delete(id int) error {
	return nil
}

func (s *Sample) GetAll() ([]schema.Todo, error) {
	todoList := []schema.Todo{
		{
			ID:       1,
			Title:    "Do dishes",
			Complete: true,
		},
		{
			ID:       2,
			Title:    "Do homework",
			Complete: false,
		},
		{
			ID:       2,
			Title:    "Twitter",
			Complete: true,
		},
	}

	return todoList, nil
}
