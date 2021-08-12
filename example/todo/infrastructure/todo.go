package infrastructure

import (
	"context"
	"errors"
	"strconv"

	"flamingo.me/graphql/example/todo/domain"
)

// TodoService service definition
type TodoService struct{}

var todos = []*domain.Todo{
	{ID: "task-0", Task: "task a"},
	{ID: "task-1", Task: "task b"},
	{ID: "task-2", Task: "task c"},
}

// Todos returns a list of mocked todos
func (ts *TodoService) Todos(_ context.Context, _ string) ([]*domain.Todo, error) {
	return todos, nil
}

// AddTodo mutation adds an entry to the list
func (ts *TodoService) AddTodo(_ context.Context, _ string, task string) (*domain.Todo, error) {
	if task == "" {
		return nil, errors.New("no todo given")
	}
	todo := &domain.Todo{
		ID:   "task-" + strconv.Itoa(len(todos)),
		Task: task,
	}
	todos = append(todos, todo)
	return todo, nil
}

// TodoDone marks a task as finished
func (ts *TodoService) TodoDone(_ context.Context, todoID string, done bool) (*domain.Todo, error) {
	for i, todo := range todos {
		if todo.ID == todoID {
			todos[i].Done = done
			return todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}
