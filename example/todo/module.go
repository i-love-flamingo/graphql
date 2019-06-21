package todo

import (
	"context"

	"flamingo.me/dingo"
	"flamingo.me/graphql"
	"flamingo.me/graphql/example/user"
	"github.com/99designs/gqlgen/codegen/config"
)

type Service struct{}

func (*Service) Schema() []byte {
	// language=graphql
	return []byte(`
type Todo {
	id: ID
	task: String!
}

extend type User {
	todos: [Todo]
}
`)
}

type Todo struct {
	ID   string
	Task string
}

func (*Service) Models() map[string]config.TypeMapEntry {
	return graphql.ModelMap{
		"Todo": Todo{},
	}.Models()
}

type Module struct{}

func (*Module) Configure(injector *dingo.Injector) {
	injector.BindMulti(new(graphql.Service)).To(new(Service))
}

func (*Module) Depends() []dingo.Module {
	return []dingo.Module{
		new(graphql.Module),
		new(user.Module),
	}
}

type TodoService struct{}

func (ts *TodoService) Todos(ctx context.Context, userid string) ([]*Todo, error) {
	return []*Todo{
		{Task: "a" + userid},
		{Task: "b" + userid},
		{Task: "c" + userid},
	}, nil
}
