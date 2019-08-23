package main

//go:generate rm -f graphql/generated.go
//go:generate go run -tags graphql main.go graphql

import (
	"bytes"
	"context"
	"io/ioutil"

	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3"
	"flamingo.me/flamingo/v3/core/requestlogger"
	"flamingo.me/flamingo/v3/framework/web"
	projectgraphql "flamingo.me/graphql/example/graphql"
	"flamingo.me/graphql/example/todo"
)

func main() {
	flamingo.App([]dingo.Module{
		new(todo.Module),
		new(projectgraphql.Module),
		new(requestlogger.Module),
		new(module),
	})
}

type module struct{}

func (*module) Configure(injector *dingo.Injector) {
	web.BindRoutes(injector, new(routes))
}

type routes struct{}

func (*routes) Routes(registry *web.RouterRegistry) {
	registry.HandleAny("static", filehandler)
}

func filehandler(ctx context.Context, r *web.Request) web.Result {
	b, _ := ioutil.ReadFile(r.Params["file"])
	return &web.Response{Status: 200, Body: bytes.NewReader(b)}
}
