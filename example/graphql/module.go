package graphql

import (
	"flamingo.me/dingo"
	"github.com/99designs/gqlgen/graphql"
)

type Module struct{}

func (*Module) Configure(injector *dingo.Injector) {
	injector.Bind(new(graphql.ExecutableSchema)).ToInstance(NewExecutableSchema(Config{Resolvers: nil}))
}

