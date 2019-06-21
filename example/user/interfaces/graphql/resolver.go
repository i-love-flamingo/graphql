package graphql

import (
	"context"

	"flamingo.me/graphql/example/user/domain"
)

type UserQueryResolver struct {
	userService domain.UserService
}

func (r *UserQueryResolver) Inject(userService domain.UserService) *UserQueryResolver {
	r.userService = userService
	return r
}

func (r *UserQueryResolver) User(ctx context.Context, id string) (*domain.User, error) {
	return r.userService.UserByID(ctx, id)
}
