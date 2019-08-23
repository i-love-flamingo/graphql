package infrastructure

import (
	"context"
	"errors"
	"strconv"

	"flamingo.me/graphql/example/todo/domain"
)

type TodoService struct{}

var todos = []*domain.Todo{
	{ID: "task-0", Task: "task a"},
	{ID: "task-1", Task: "task b"},
	{ID: "task-2", Task: "task c"},
}

func (ts *TodoService) Todos(ctx context.Context, userid string) ([]*domain.Todo, error) {
	return todos, nil
}

func (ts *TodoService) AddTodo(ctx context.Context, userid string, task string) (*domain.Todo, error) {
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
