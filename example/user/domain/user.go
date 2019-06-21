package domain

import "context"

type User struct {
	Name string
}

type UserService interface {
	UserByID(ctx context.Context, id string) (*User, error)
}
