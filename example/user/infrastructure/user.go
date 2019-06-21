package infrastructure

import (
	"context"

	"flamingo.me/graphql/example/user/domain"
)

type UserServiceImpl struct{}

func (us *UserServiceImpl) UserByID(ctx context.Context, id string) (*domain.User, error) {
	return &domain.User{
		Name: "user-" + id,
	}, nil
}
