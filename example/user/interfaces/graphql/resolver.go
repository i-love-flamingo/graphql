package graphql

import (
	"context"

	"flamingo.me/graphql/example/user/domain"
)

// UserQueryResolver resolver for the user service
type UserQueryResolver struct {
	userService domain.UserService
}

// Inject dependencies
func (r *UserQueryResolver) Inject(userService domain.UserService) *UserQueryResolver {
	r.userService = userService
	return r
}

// User getter
func (r *UserQueryResolver) User(ctx context.Context, id string) (*domain.User, error) {
	return r.userService.UserByID(ctx, id)
}
