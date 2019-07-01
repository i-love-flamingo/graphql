package graphql

import (
	"flamingo.me/graphql/templates"
	"io/ioutil"
	"log"
	"os"
	"path"
	"reflect"
	"strings"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/spf13/cobra"
)

const (
	basePath       = "graphql"
	schemaPath     = "graphql/schema"
	baseSchemaPath = "graphql/schema/schema.graphql"
	resolverPath   = "graphql/resolver.go"
	execPath       = "graphql/generated.go"
	modelPath      = "graphql/models_gen.go"
)

func command(
	services []Service,
) *cobra.Command {
	return &cobra.Command{
		Use: "graphql flamingo module",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := config.DefaultConfig()
			cfg.SchemaFilename = []string{baseSchemaPath}
			cfg.Models = make(map[string]config.TypeMapEntry)

			if err := os.MkdirAll(schemaPath, 0755); err != nil && !os.IsExist(err) {
				return err
			}

			if err := ioutil.WriteFile(baseSchemaPath, templates.MustAsset("schema.graphql"), 0644); err != nil {
				return err
			}

			if err := ioutil.WriteFile(path.Join(basePath, "module.go"), templates.MustAsset("module.go.tpl"), 0644); err != nil {
				return err
			}

			if _, err := os.Stat(resolverPath); os.IsNotExist(err) {
				if err := ioutil.WriteFile(resolverPath, templates.MustAsset("resolver.go.tpl"), 644); err != nil {
					return err
				}
			}

			for _, service := range services {
				rt := reflect.TypeOf(service).Elem()
				fname := strings.Replace(rt.PkgPath(), "/", "_", -1) + "-" + rt.Name() + ".graphql"
				fpath := path.Join(schemaPath, fname)

				log.Printf("Writing %s", fname)
				if err := ioutil.WriteFile(fpath, service.Schema(), 0644); err != nil {
					return err
				}
				cfg.SchemaFilename = append(cfg.SchemaFilename, fpath)

				// merge models into config models
				for k, v := range service.Models() {
					cfg.Models[k] = v
				}
			}

			float := cfg.Models["Float"]
			float.Model = append(float.Model, "flamingo.me/graphql.Float", "github.com/99designs/gqlgen/graphql.Float")
			cfg.Models["Float"] = float

			cfg.Model = config.PackageConfig{Filename: modelPath}
			cfg.Exec = config.PackageConfig{Filename: execPath}

			return api.Generate(cfg)
		},
	}
}
