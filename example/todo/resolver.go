package todo

import (
	"context"

	"flamingo.me/graphql/example/user/domain"
)

type TodoUserResolver struct {
	todosBackend *TodoService
}

func (r *TodoUserResolver) Inject(todosBackend *TodoService) *TodoUserResolver {
	r.todosBackend = todosBackend
	return r
}

func (r *TodoUserResolver) Todos(ctx context.Context, obj *domain.User) ([]*Todo, error) {
	return r.todosBackend.Todos(ctx, obj.Name)
}
