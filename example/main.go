package main

//go:generate rm -f graphql/generated.go
//go:generate go run -tags graphql main.go graphql

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3"
	"flamingo.me/flamingo/v3/core/requestlogger"
	"flamingo.me/graphql/example/todo"
)

func main() {
	flamingo.App([]dingo.Module{
		new(todo.Module),
		new(requestlogger.Module),
	})
}
