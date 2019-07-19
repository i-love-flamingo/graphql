package main

//go:generate rm -f graphql/generated.go
//go:generate go run -tags graphql main.go graphql

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3"
	"flamingo.me/graphql"
	projectgraphql "flamingo.me/graphql/example/graphql"
	"flamingo.me/graphql/example/todo"
	"flamingo.me/graphql/example/user"
)

func main() {
	flamingo.App([]dingo.Module{
		new(graphql.Module), // implicit via dependency
		new(user.Module),    // implicit via dependency
		new(todo.Module),    // implicit via dependency
		new(projectgraphql.Module),
	})
}
