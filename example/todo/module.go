package todo

import (
	"flamingo.me/dingo"
	"flamingo.me/graphql"
	"flamingo.me/graphql/example/todo/domain"
	"flamingo.me/graphql/example/user"
	"github.com/99designs/gqlgen/codegen/config"
)

type Service struct{}

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

func (*Service) Models() map[string]config.TypeMapEntry {
	return graphql.ModelMap{
		"Todo": domain.Todo{},
	}.Models()
}

type Module struct{}

func (Module) Configure(injector *dingo.Injector) {
	injector.BindMulti(new(graphql.Service)).To(new(Service))
}

func (*Module) Depends() []dingo.Module {
	return []dingo.Module{
		new(user.Module),
	}
}
