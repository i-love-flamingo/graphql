package todo

import (
	"context"
	"math/big"

	"flamingo.me/dingo"
	"flamingo.me/graphql"
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

type Todo struct {
	ID      string
	Task    string
	Points  *big.Float
	Points2 float64
}

type Task interface {
	A() string
}

func (*Todo) A() string {
	return "a"
}

func (*Service) Models() map[string]config.TypeMapEntry {
	return graphql.ModelMap{
		"Todo": Todo{},
		"Task": new(Task),
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
		{Task: "a" + userid, Points: big.NewFloat(12.34), Points2: 22.34},
		{Task: "b" + userid},
		{Task: "c" + userid},
	}, nil
}

type TodoResolver struct{}

func (*TodoResolver) B(ctx context.Context, obj *Todo) (*string, error) {
	s := "B"
	return &s, nil
}
