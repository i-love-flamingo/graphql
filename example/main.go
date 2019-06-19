package main

//go:generate go run main.go graphql

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3"
	projectgraphql "flamingo.me/graphql/example/graphql"
	"flamingo.me/graphql/example/todo"
)

func main() {
	flamingo.App([]dingo.Module{
		// new(graphql.Module),  // implicit via dependency
		// new(user.Module),  // implicit via dependency
		new(todo.Module),
		new(projectgraphql.Module),
	})
}
