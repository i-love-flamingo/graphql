package user

import (
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
`)
}

type User struct {
	Name string
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
