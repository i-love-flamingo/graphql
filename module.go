package graphql

import (
	"context"
	"time"

	"flamingo.me/dingo"
	flamingoConfig "flamingo.me/flamingo/v3/framework/config"
	"flamingo.me/flamingo/v3/framework/web"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/spf13/cobra"
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
	uploadMaxSize        int64
	operationMiddlewares []graphql.OperationMiddleware
}

// Inject executable schema
func (r *routes) Inject(
	reverseRouter web.ReverseRouter,
	config *struct {
		Exec                 graphql.ExecutableSchema      `inject:",optional"`
		OperationMiddlewares []graphql.OperationMiddleware `inject:",optional"`
		Origins              flamingoConfig.Slice          `inject:"config:graphql.cors.origins"`
		IntrospectionEnabled bool                          `inject:"config:graphql.introspectionEnabled,optional"`
		UploadMaxSize        int64                         `inject:"config:graphql.multipartForm.uploadMaxSize,optional"`
	},
) {
	r.reverseRouter = reverseRouter
	if config != nil {
		r.exec = config.Exec
		r.origins = config.Origins
		r.introspectionEnabled = config.IntrospectionEnabled
		r.uploadMaxSize = config.UploadMaxSize
		r.operationMiddlewares = config.OperationMiddlewares
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
	gqlHandler := func(es graphql.ExecutableSchema) *handler.Server {
		srv := handler.New(es)

		srv.AddTransport(transport.Websocket{
			KeepAlivePingInterval: 10 * time.Second,
		})
		srv.AddTransport(transport.Options{})
		srv.AddTransport(transport.GET{})
		srv.AddTransport(transport.POST{})
		srv.AddTransport(transport.MultipartForm{
			MaxUploadSize: r.uploadMaxSize,
		})
		srv.SetQueryCache(lru.New(1000))

		srv.Use(extension.Introspection{})
		srv.Use(extension.AutomaticPersistedQuery{
			Cache: lru.New(100),
		})

		return srv
	}(r.exec)

	gqlHandler.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		graphql.GetOperationContext(ctx).DisableIntrospection = !r.introspectionEnabled
		return next(ctx)
	})

	for _, middleware := range r.operationMiddlewares {
		gqlHandler.AroundOperations(middleware)
	}

	registry.MustRoute("/graphql", "graphql")
	registry.HandleOptions("graphql", web.WrapHTTPHandler(corsHandler.preflightHandler()))
	registry.HandleAny("graphql", wrapGqlHandler(corsHandler.gqlMiddleware(gqlHandler)))

	registry.MustRoute("/graphql-console", "graphql.console")
	u, _ := r.reverseRouter.Relative("graphql", nil)
	registry.HandleAny("graphql.console", web.WrapHTTPHandler(playground.Handler("Flamingo GraphQL Console", u.String())))
}

// CueConfig for the module
func (m *Module) CueConfig() string {
	return `
graphql: {
	introspectionEnabled: bool | *false
	multipartForm: {
		uploadMaxSize: (int | *1.5M) & > 0
	}
}
`
}
