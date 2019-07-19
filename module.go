package graphql

//go:generate go run github.com/go-bindata/go-bindata/go-bindata -o templates/fs.go -pkg templates -prefix templates/ templates/*.tpl templates/schema.graphql

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3/framework/web"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/handler"
	"github.com/spf13/cobra"
)

// Service defines the bind point for graphql services
type Service interface {
	Schema() []byte
	Models() map[string]config.TypeMapEntry
}

// Module defines the graphql entry point and binds the graphql command and routes
type Module struct{}

// Configure sets up dingo
func (*Module) Configure(injector *dingo.Injector) {
	injector.BindMulti(new(cobra.Command)).ToProvider(command)

	web.BindRoutes(injector, new(routes))
}

type routes struct {
	exec          graphql.ExecutableSchema
	reverseRouter web.ReverseRouter
}

// Inject executable schema
func (r *routes) Inject(config *struct {
	Exec graphql.ExecutableSchema `inject:",optional"`
}, reverseRouter web.ReverseRouter) {
	r.exec = config.Exec
	r.reverseRouter = reverseRouter
}

// Routes definition for flamingo router
func (r *routes) Routes(registry *web.RouterRegistry) {
	if r.exec == nil {
		panic("Please register a schema/module before running the server!")
	}

	registry.Route("/graphql", "graphql")
	registry.HandleAny("graphql", web.WrapHTTPHandler(handler.GraphQL(r.exec)))

	registry.Route("/graphql-console", "graphql.console")
	u, _ := r.reverseRouter.Relative("graphql", nil)
	registry.HandleAny("graphql.console", web.WrapHTTPHandler(handler.Playground("Flamingo GraphQL Console", u.String())))
}
