package user

import (
	"context"

	"flamingo.me/dingo"
	"flamingo.me/graphql"
	"github.com/99designs/gqlgen/codegen/config"
)

type Service struct{}

func (*Service) Schema() []byte {
	// language=graphql
	return []byte(`
type User {
	name: String!
}

extend type Query {
	User(id: String!): User
}
`)
}

type User struct {
	Name string
}

type UserService interface {
	UserByID(ctx context.Context, id string) (*User, error)
}

func (*Service) Models() map[string]config.TypeMapEntry {
	return graphql.ModelMap{
		"User": User{},
	}.Models()
}

type Module struct{}

func (*Module) Configure(injector *dingo.Injector) {
	injector.BindMulti(new(graphql.Service)).To(new(Service))
}
