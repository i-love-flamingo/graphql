package main

//go:generate go run main.go graphql

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
		new(graphql.Module),
		new(user.Module),
		new(todo.Module),
		new(projectgraphql.Module),
	})
}
