package todo

import (
	"context"

	"flamingo.me/graphql/example/todo/domain"
	"flamingo.me/graphql/example/todo/infrastructure"
	userDomain "flamingo.me/graphql/example/user/domain"
)

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

type TodoMutationResolver struct {
	resolver *TodoUserResolver
}

func (r *TodoMutationResolver) Inject(resolver *TodoUserResolver) *TodoMutationResolver {
	r.resolver = resolver
	return r
}

func (r *TodoMutationResolver) TodoAdd(ctx context.Context, user string, task string) (*domain.Todo, error) {
	return r.resolver.todosBackend.AddTodo(ctx, user, task)
}

func (r *TodoMutationResolver) TodoDone(ctx context.Context, todo string, done bool) (*domain.Todo, error) {
	return r.resolver.todosBackend.TodoDone(ctx, todo, done)
}
