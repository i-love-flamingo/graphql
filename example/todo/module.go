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
interface Task {
	a: String
}

type Todo implements Task {
	id: ID
	task: String!
	points: Float
	points2: Float

	a: String
	b: String
}

extend interface Task {
	b: String
}

extend type User {
	todos: [Todo]
}
`)
}

func (*Service) Models() map[string]config.TypeMapEntry {
	return graphql.ModelMap{
		"Todo": domain.Todo{},
		"Task": new(domain.Task),
	}.Models()
}

type Module struct{}

func (Module) Configure(injector *dingo.Injector) {
	injector.BindMulti(new(graphql.Service)).To(new(Service))
}

func (*Module) Depends() []dingo.Module {
	return []dingo.Module{
		new(user.Module),
		new(graphql.Module),
	}
}
