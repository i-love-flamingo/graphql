package todo

import (
	"context"

	"flamingo.me/graphql/example/todo/domain"
	"flamingo.me/graphql/example/todo/infrastructure"
	userDomain "flamingo.me/graphql/example/user/domain"
)

type TodoResolver struct{}

func (*TodoResolver) B(ctx context.Context, obj *domain.Todo) (*string, error) {
	s := "B"
	return &s, nil
}

type TodoUserResolver struct {
	todosBackend *infrastructure.TodoService
}

func (r *TodoUserResolver) Inject(todosBackend *infrastructure.TodoService) *TodoUserResolver {
	r.todosBackend = todosBackend
	return r
}

func (r *TodoUserResolver) Todos(ctx context.Context, obj *userDomain.User) ([]*domain.Todo, error) {
	return r.todosBackend.Todos(ctx, obj.Name)
}
