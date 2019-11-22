package graphql // import "flamingo.me/graphql"

//go:generate go run github.com/go-bindata/go-bindata/go-bindata -nometadata -prefix templates/ -o templates/fs.go -pkg templates -ignore=fs.go templates/

import (
	"flamingo.me/dingo"
	flamingoConfig "flamingo.me/flamingo/v3/framework/config"
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
	exec                 graphql.ExecutableSchema
	reverseRouter        web.ReverseRouter
	whitelist            flamingoConfig.Slice
	introspectionEnabled bool
}

// Inject executable schema
func (r *routes) Inject(
	reverseRouter web.ReverseRouter,
	config *struct {
		Exec                 graphql.ExecutableSchema `inject:",optional"`
		Whitelist            flamingoConfig.Slice     `inject:"config:graphql.cors.whitelist"`
		IntrospectionEnabled bool                     `inject:"config:graphql.introspectionEnabled,optional"`
	},
) {
	r.reverseRouter = reverseRouter
	if config != nil {
		r.exec = config.Exec
		r.whitelist = config.Whitelist
		r.introspectionEnabled = config.IntrospectionEnabled
	}
}

// Routes definition for flamingo router
func (r *routes) Routes(registry *web.RouterRegistry) {
	if r.exec == nil {
		panic("Please register/generate a schema module before running the server!")
	}

	var whitelist []string
	r.whitelist.MapInto(&whitelist)

	corsHandler := corsHandler{whitelist: whitelist}

	registry.Route("/graphql", "graphql")
	registry.HandleOptions("graphql", web.WrapHTTPHandler(corsHandler.preflightHandler()))
	registry.HandleAny("graphql",
		web.WrapHTTPHandler(
			corsHandler.gqlMiddleware(
				handler.GraphQL(
					r.exec,
					handler.IntrospectionEnabled(r.introspectionEnabled),
				),
			),
		),
	)

	registry.Route("/graphql-console", "graphql.console")
	u, _ := r.reverseRouter.Relative("graphql", nil)
	registry.HandleAny("graphql.console", web.WrapHTTPHandler(handler.Playground("Flamingo GraphQL Console", u.String())))
}

// DefaultConfig for this module
func (m *Module) DefaultConfig() flamingoConfig.Map {
	return flamingoConfig.Map{
		"graphql.introspectionEnabled": false,
	}
}
