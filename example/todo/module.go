package todo

import (
	"flamingo.me/dingo"
	"flamingo.me/graphql"
	"flamingo.me/graphql/example/todo/domain"
	"flamingo.me/graphql/example/user"
	"github.com/99designs/gqlgen/codegen/config"
)

// Service for the todo graphql service
type Service struct{}

// Schema getter for graphql
func (*Service) Schema() []byte {
	// language=graphql
	return []byte(`
type Todo {
	id: ID!
	task: String!
	done: Boolean!
}

extend type User {
	todos: [Todo]
}

extend type Mutation {
	TodoAdd(user: ID!, task: String!): Todo
	TodoDone(todo: ID!, done: Boolean!): Todo
}
`)
}

// Models mapping between graphql and go
func (*Service) Models() map[string]config.TypeMapEntry {
	return graphql.ModelMap{
		"Todo": domain.Todo{},
	}.Models()
}

// Module struct for the todo module
type Module struct{}

// Configure registers the service graphql service
func (Module) Configure(injector *dingo.Injector) {
	injector.BindMulti(new(graphql.Service)).To(new(Service))
}

// Depends marks for the user Module
func (*Module) Depends() []dingo.Module {
	return []dingo.Module{
		new(user.Module),
	}
}
