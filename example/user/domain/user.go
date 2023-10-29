package domain

import "context"

// User contains the basic user information
type User struct {
	Name       string
	Nicknames  []string
	Attributes Attributes
}

// Attributes for random user attributes
type Attributes map[string]interface{}

// UserService allows to retrieve a user
type UserService interface {
	UserByID(ctx context.Context, id string) (*User, error)
}

// Get returns an attribute by its key
func (a Attributes) Get(key string) string {
	if s, ok := a[key].(string); ok {
		return s
	}

	return ""
}

// Keys lists all attribute keys
func (a Attributes) Keys() []string {
	keys := make([]string, len(a))
	i := 0

	for k := range a {
		keys[i] = k
		i++
	}

	return keys
}
