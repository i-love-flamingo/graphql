package graphql

import (
	"flamingo.me/dingo"
	"flamingo.me/graphql"
	"flamingo.me/graphql/example/todo"
	"flamingo.me/graphql/example/user"
)

func (*Module) Depends() []dingo.Module {
	return []dingo.Module{
		new(user.Module),
		new(todo.Module),
		new(graphql.Module),
	}
}
