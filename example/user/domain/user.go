package domain

import "context"

type User struct {
	Name      string
	Nicknames []string
}

type UserService interface {
	UserByID(ctx context.Context, id string) (*User, error)
}
