package graphql

import (
	"context"

	"flamingo.me/graphql"
	"flamingo.me/graphql/example/todo"
	"flamingo.me/graphql/example/user"
)

type rootResolver struct {
	queryResolver *queryResolver
	userResolver  *userResolver
}

var _ ResolverRoot = new(rootResolver)

func (r *rootResolver) Query() QueryResolver {
	return r.queryResolver
}

func (r *rootResolver) User() UserResolver {
	return r.userResolver
}

type queryResolver struct {
	userService user.UserService
}

var _ QueryResolver = new(queryResolver)

func (r *queryResolver) Inject(userService user.UserService) {
	r.userService = userService
}

func (*queryResolver) Flamingo(ctx context.Context) (*string, error) {
	return graphql.FlamingoResolver(ctx)
}

func (r *queryResolver) User(ctx context.Context, id string) (*user.User, error) {
	return r.userService.UserByID(ctx, id)
}

type userResolver struct{}

func (*userResolver) Todos(ctx context.Context, obj *user.User) ([]*todo.Todo, error) {
	return nil, nil
}
