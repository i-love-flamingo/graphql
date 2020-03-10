package todo

import (
	"context"

	"flamingo.me/graphql/example/todo/domain"
	"flamingo.me/graphql/example/todo/infrastructure"
	userDomain "flamingo.me/graphql/example/user/domain"
)

// UserResolver mapper between graphql and the todos backend
type UserResolver struct {
	todosBackend *infrastructure.TodoService
}

// Inject dependencies
func (r *UserResolver) Inject(todosBackend *infrastructure.TodoService) *UserResolver {
	r.todosBackend = todosBackend
	return r
}

// Todos getter
func (r *UserResolver) Todos(ctx context.Context, obj *userDomain.User) ([]*domain.Todo, error) {
	return r.todosBackend.Todos(ctx, obj.Name)
}

// MutationResolver maps mutations
type MutationResolver struct {
	resolver *UserResolver
}

// Inject dependencies
func (r *MutationResolver) Inject(resolver *UserResolver) *MutationResolver {
	r.resolver = resolver
	return r
}

// TodoAdd mutation
func (r *MutationResolver) TodoAdd(ctx context.Context, user string, task string) (*domain.Todo, error) {
	return r.resolver.todosBackend.AddTodo(ctx, user, task)
}

// TodoDone mutation
func (r *MutationResolver) TodoDone(ctx context.Context, todo string, done bool) (*domain.Todo, error) {
	return r.resolver.todosBackend.TodoDone(ctx, todo, done)
}
