package graphql

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/99designs/gqlgen/graphql"

	"github.com/99designs/gqlgen/handler"

	"flamingo.me/flamingo/v3/framework/web"

	"github.com/99designs/gqlgen/api"

	"flamingo.me/dingo"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/spf13/cobra"
)

type Service interface {
	Schema() []byte
	Models() map[string]config.TypeMapEntry
}

type ModelMap map[string]interface{}

func (m ModelMap) Models() map[string]config.TypeMapEntry {
	res := make(map[string]config.TypeMapEntry, len(m))
	for k, v := range m {
		v := reflect.TypeOf(v)
		res[k] = config.TypeMapEntry{
			Model: []string{v.PkgPath() + "." + v.Name()},
		}
	}
	return res
}

type Module struct{}

func (*Module) Configure(injector *dingo.Injector) {
	injector.BindMulti(new(cobra.Command)).ToProvider(func(
		services []Service,
	) *cobra.Command {
		return &cobra.Command{
			Use: "graphql flamingo module",
			Run: func(cmd *cobra.Command, args []string) {
				cfg := config.DefaultConfig()
				cfg.SchemaFilename = []string{"graphql/schema.graphql"}
				cfg.Models = make(map[string]config.TypeMapEntry)

				os.Mkdir("graphql", 0755)

				ioutil.WriteFile("graphql/schema.graphql", []byte(`type Query { flamingo: String }`), 0644)

				// language=go
				ioutil.WriteFile("graphql/module.go", []byte(`package graphql

import (
	"flamingo.me/dingo"
	"github.com/99designs/gqlgen/graphql"
)

type Module struct{}

func (*Module) Configure(injector *dingo.Injector) {
	injector.Bind(new(graphql.ExecutableSchema)).ToInstance(NewExecutableSchema(Config{Resolvers: nil}))
}

`), 0644)

				for _, service := range services {
					rt := reflect.TypeOf(service).Elem()
					fname := strings.Replace(rt.PkgPath(), "/", "_", -1) + "-" + rt.Name() + ".graphql"
					log.Println(fname)
					ioutil.WriteFile("graphql/"+fname, service.Schema(), 0644)
					cfg.SchemaFilename = append(cfg.SchemaFilename, "graphql/"+fname)

					// merge models into config models
					for k, v := range service.Models() {
						cfg.Models[k] = v
					}
				}

				cfg.Model = config.PackageConfig{Filename: "graphql/models_gen.go"}
				cfg.Exec = config.PackageConfig{Filename: "graphql/generated.go"}

				fmt.Println(api.Generate(cfg))
			},
		}
	})

	web.BindRoutes(injector, new(routes))
}

type routes struct {
	exec graphql.ExecutableSchema
}

func (r *routes) Inject(exec graphql.ExecutableSchema) {
	r.exec = exec
}

func (r *routes) Routes(registry *web.RouterRegistry) {
	registry.Route("/graphql", "graphql")
	registry.HandleAny("graphql", web.WrapHTTPHandler(handler.GraphQL(r.exec)))

	registry.Route("/graphql-console", "graphql.console")
	registry.HandleAny("graphql.console", web.WrapHTTPHandler(handler.Playground("Flamingo GraphQL Console", "/graphql")))
}
