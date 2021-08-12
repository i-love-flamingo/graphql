package graphql

//go:generate go run github.com/go-bindata/go-bindata/v3/go-bindata -nometadata -prefix templates/ -o templates/fs.go -pkg templates -ignore=fs.go templates/

import (
	"context"

	"flamingo.me/dingo"
	flamingoConfig "flamingo.me/flamingo/v3/framework/config"
	"flamingo.me/flamingo/v3/framework/web"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/spf13/cobra"

	// keep dependency from cleaning up
	_ "github.com/go-bindata/go-bindata/v3"
)

// Service defines the interface for graphql services
// The Schema returns the GraphQL Schema definition
// The Types configure the GraphQL type mapping and resolution
type Service interface {
	Schema() []byte
	Types(types *Types)
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
	origins              flamingoConfig.Slice
	introspectionEnabled bool
}

// Inject executable schema
func (r *routes) Inject(
	reverseRouter web.ReverseRouter,
	config *struct {
		Exec                 graphql.ExecutableSchema `inject:",optional"`
		Origins              flamingoConfig.Slice     `inject:"config:graphql.cors.origins"`
		IntrospectionEnabled bool                     `inject:"config:graphql.introspectionEnabled,optional"`
	},
) {
	r.reverseRouter = reverseRouter
	if config != nil {
		r.exec = config.Exec
		r.origins = config.Origins
		r.introspectionEnabled = config.IntrospectionEnabled
	}
}

// Routes definition for flamingo router
func (r *routes) Routes(registry *web.RouterRegistry) {
	if r.exec == nil {
		panic("Please register/generate a schema module before running the server!")
	}

	var origins []string
	err := r.origins.MapInto(&origins)
	if err != nil {
		panic(err)
	}

	corsHandler := corsHandler{origins: origins}
	gqlHandler := handler.NewDefaultServer(r.exec)
	gqlHandler.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		graphql.GetOperationContext(ctx).DisableIntrospection = !r.introspectionEnabled
		return next(ctx)
	})

	registry.MustRoute("/graphql", "graphql")
	registry.HandleOptions("graphql", web.WrapHTTPHandler(corsHandler.preflightHandler()))
	registry.HandleAny("graphql", web.WrapHTTPHandler(corsHandler.gqlMiddleware(gqlHandler)))

	registry.MustRoute("/graphql-console", "graphql.console")
	u, _ := r.reverseRouter.Relative("graphql", nil)
	registry.HandleAny("graphql.console", web.WrapHTTPHandler(playground.Handler("Flamingo GraphQL Console", u.String())))
}

// DefaultConfig for this module
func (m *Module) DefaultConfig() flamingoConfig.Map {
	return flamingoConfig.Map{
		"graphql.introspectionEnabled": false,
	}
}
