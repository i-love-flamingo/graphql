package domain

import "context"

// User contains the basic user information
type User struct {
	Name      string
	Nicknames []string
}

// UserService allows to retrieve a user
type UserService interface {
	UserByID(ctx context.Context, id string) (*User, error)
}
