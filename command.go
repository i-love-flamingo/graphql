package graphql

import (
	// enable go:embed
	_ "embed"
	"fmt"
	"log"
	"os"
	"path"
	"reflect"
	"strings"
	"syscall"
	"text/template"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen"
	"github.com/99designs/gqlgen/codegen/config"
	gqltemplates "github.com/99designs/gqlgen/codegen/templates"
	plugin2 "github.com/99designs/gqlgen/plugin"
	"github.com/spf13/cobra"
	"github.com/vektah/gqlparser/v2/ast"
)

const (
	basePath       = "graphql"
	schemaBasePath = "graphql/schema"
)

var skipGoModTidy = true

//go:embed templates/schema.graphql
var schema []byte

//go:embed templates/module.go.tpl
var module []byte

//go:embed templates/emptymodule.go.tpl
var emptyModule []byte

func command(
	services []Service,
) *cobra.Command {
	return &cobra.Command{
		Use:     "graphql",
		Short:   "(Re-)Generate graphql module",
		Aliases: []string{"g"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return Generate(services, basePath, schemaBasePath)
		},
	}
}

// Generate runs the graphql generation for the defined services
// Usually you won't need this, and it's intended for infrastructure implementations.
func Generate(services []Service, basePath string, schemaBasePath string) error {
	schemaPath := path.Join(schemaBasePath, "schema.graphql")

	cfg := config.DefaultConfig()
	err := config.CompleteConfig(cfg)
	if err != nil {
		return err
	}

	cfg.SchemaFilename = []string{schemaPath}
	cfg.Models = make(map[string]config.TypeMapEntry)

	if err := os.MkdirAll(schemaBasePath, 0755); err != nil && !os.IsExist(err) {
		return fmt.Errorf("mkdir %q failed: %w", schemaBasePath, err)
	}

	if err := os.WriteFile(schemaPath, schema, 0644); err != nil {
		return fmt.Errorf("writefile %q failed: %w", schemaPath, err)
	}

	if err := os.WriteFile(path.Join(basePath, "module.go"), module, 0644); err != nil {
		return fmt.Errorf("writefile %q/module.go failed: %w", basePath, err)
	}

	if err := os.WriteFile(path.Join(basePath, "emptymodule.go"), emptyModule, 0644); err != nil {
		return fmt.Errorf("writefile %q/emptymodule.go failed: %w", basePath, err)
	}

	types := new(Types)

	types.Resolve("Query", "flamingo", FlamingoQueryResolver{}, "Flamingo")
	types.Resolve("Mutation", "flamingo", FlamingoQueryResolver{}, "Flamingo")

	for _, service := range services {
		rt := reflect.TypeOf(service).Elem()
		fname := strings.ReplaceAll(rt.PkgPath(), "/", "_") + "-" + rt.Name() + ".graphql"
		fpath := path.Join(schemaBasePath, fname)

		log.Printf("Writing %s", fname)
		if err := os.WriteFile(fpath, service.Schema(), 0644); err != nil {
			return fmt.Errorf("writefile %q failed: %w", fpath, err)
		}
		cfg.SchemaFilename = append(cfg.SchemaFilename, fpath)

		service.Types(types)
	}

	// merge models into config models
	for graphqlObject, goType := range types.names {
		cfg.Models[graphqlObject] = config.TypeMapEntry{Model: []string{goType}}
	}

	for graphqlObject, fields := range types.fields {
		for graphqlField, goType := range fields {
			model := cfg.Models[graphqlObject]
			if cfg.Models[graphqlObject].Fields == nil {
				model.Fields = make(map[string]config.TypeMapField)
			}
			model.Fields[graphqlField] = config.TypeMapField{FieldName: goType}
			cfg.Models[graphqlObject] = model
		}
	}

	for graphqlObject, resolver := range types.resolver {
		for graphqlField := range resolver {
			model := cfg.Models[graphqlObject]
			if cfg.Models[graphqlObject].Fields == nil {
				model.Fields = make(map[string]config.TypeMapField)
			}
			model.Fields[graphqlField] = config.TypeMapField{Resolver: true}
			cfg.Models[graphqlObject] = model
		}
	}

	float := cfg.Models["Float"]
	float.Model = append(float.Model, "flamingo.me/graphql.Float", "github.com/99designs/gqlgen/graphql.Float")
	cfg.Models["Float"] = float

	cfg.Models["Date"] = config.TypeMapEntry{
		Model: []string{"flamingo.me/graphql.Date"},
	}

	cfg.Model = config.PackageConfig{}
	cfg.Exec = config.ExecConfig{Filename: path.Join(basePath, "generated.go")}
	cfg.SkipModTidy = skipGoModTidy

	for _, filename := range cfg.SchemaFilename {
		schemaRaw, err := os.ReadFile(filename)
		if err != nil {
			return fmt.Errorf("unable to open schema: %w", err)
		}

		cfg.Sources = append(cfg.Sources, &ast.Source{Name: filename, Input: string(schemaRaw)})
	}

	if err := api.Generate(cfg, api.AddPlugin(&plugin{types: types})); err != nil {
		return fmt.Errorf("gqlgen/api.Generate failed: %w", err)
	}
	return nil
}

var _ plugin2.CodeGenerator = &plugin{}
var _ plugin2.ConfigMutator = &plugin{}
var _ plugin2.Plugin = &plugin{}

type plugin struct {
	types *Types
}

func (*plugin) Name() string {
	return "flamingo.me/graphql"
}

func (m *plugin) MutateConfig(_ *config.Config) error {
	_ = syscall.Unlink("graphql/resolver.go")
	return nil
}

func (m *plugin) GenerateCode(data *codegen.Data) error {
	// Drop directives that can't be resolved, this was needed to ignore the `@defer` directive not used by us
	for name := range data.AllDirectives {
		if pkg := m.types.directives["@"+name][0]; pkg == "" {
			delete(data.AllDirectives, name)
		}
	}

	errored := false
	defer func() {
		if errored {
			panic("code generation failed")
		}
	}()
	return gqltemplates.Render(gqltemplates.Options{
		PackageName: "graphql",
		Filename:    "graphql/resolver.go",
		Data: &resolverBuild{
			Data:     data,
			TypeName: "rootResolver",
		},
		GeneratedHeader: true,
		Packages:        data.Config.Packages,
		Funcs: template.FuncMap{
			"gpkg": func(from, to string, field codegen.Field) string {
				if m.types.resolver[from][to][0] == "" {
					fmt.Printf(
						"\nmissing resolver for %q.%q:\n\tfunc (r *%sResolver) %s%s\n\n\ttypes.Resolve(\"%s\", \"%s\", %sResolver{}, \"%s\")\n\n",
						from, to,
						gqltemplates.UcFirst(from), gqltemplates.UcFirst(to),
						field.ShortResolverDeclaration(),
						from, to,
						gqltemplates.UcFirst(from), gqltemplates.UcFirst(to),
					)
					errored = true
				}
				return m.types.resolver[from][to][0]
			},
			"gtyp": func(from, to string) string {
				return m.types.resolver[from][to][1]
			},
			"gmethod": func(from, to string) string {
				return m.types.resolver[from][to][2]
			},
			"gdpkg": func(name string) string {
				return m.types.directives["@"+name][0]
			},
			"gdtyp": func(name string) string {
				return m.types.directives["@"+name][1]
			},
			"gdmethod": func(name string) string {
				return m.types.directives["@"+name][2]
			},
		},
		PackageDoc: "//+build !graphql\n",
		Template: `
{{ reserveImport "context"  }}
{{ reserveImport "fmt"  }}
{{ reserveImport "io"  }}
{{ reserveImport "strconv"  }}
{{ reserveImport "time"  }}
{{ reserveImport "sync"  }}
{{ reserveImport "errors"  }}
{{ reserveImport "bytes"  }}
{{ reserveImport "log"  }}

{{ reserveImport "github.com/vektah/gqlparser/v2" }}
{{ reserveImport "github.com/vektah/gqlparser/v2/ast" }}
{{ reserveImport "github.com/99designs/gqlgen/graphql" }}
{{ reserveImport "github.com/99designs/gqlgen/graphql/introspection" }}

{{ $root := . }}

var _ ResolverRoot = new({{$root.TypeName}})

type {{$root.TypeName}} struct {
	{{ range $object := .Objects }}
		{{- if $object.HasResolvers }}
			{{lcFirst $root.TypeName}}{{$object.Name}} *{{lcFirst $root.TypeName}}{{$object.Name}}
		{{- end }}
	{{- end }}
	{{ range $directive := .AllDirectives }}
		{{$directive.Name}}Resolver *{{lookupImport (gdpkg $directive.Name)}}.{{gdtyp $directive.Name}}
	{{- end }}
}

func (r *{{$root.TypeName}}) Inject (
	{{- range $object := .Objects }}
		{{- if $object.HasResolvers }}
			{{lcFirst $root.TypeName}}{{$object.Name}} *{{lcFirst $root.TypeName}}{{$object.Name}},
		{{- end }}
	{{- end }}
	{{ range $directive := .AllDirectives }}
		{{$directive.Name}}Resolver *{{lookupImport (gdpkg $directive.Name)}}.{{gdtyp $directive.Name}},
	{{- end }}
) {
	{{- range $object := .Objects }}
		{{- if $object.HasResolvers }}
			r.{{lcFirst $root.TypeName}}{{$object.Name}} = {{lcFirst $root.TypeName}}{{$object.Name}}
		{{- end }}
	{{- end }}
	{{ range $directive := .AllDirectives }}
		r.{{$directive.Name}}Resolver = {{$directive.Name}}Resolver
	{{- end }}
}

func (r *{{$root.TypeName}}) directives() DirectiveRoot {
	return DirectiveRoot{
		{{ range $directive := .AllDirectives -}}
			{{ucFirst $directive.Name}}: r.{{$directive.Name}}Resolver.{{gdmethod $directive.Name}},
		{{end}}
	}
}

{{ range $object := .Objects -}}
	{{- if $object.HasResolvers -}}
		func (r *{{$.TypeName}}) {{$object.Name}}() {{ $object.ResolverInterface | ref }} {
			return r.{{lcFirst $root.TypeName}}{{$object.Name}}
		}
	{{ end -}}
{{ end }}

{{ range $object := .Objects -}}
	{{- if $object.HasResolvers -}}
		type {{lcFirst $root.TypeName}}{{$object.Name}} struct {
			{{range $field := $object.Fields }}
				{{- if $field.IsResolver }}
					{{- print}}resolve{{$field.GoFieldName}} func{{ $field.ShortResolverDeclaration }}
				{{ end }}
			{{- end }}
		}

		func (r *{{lcFirst $root.TypeName}}{{$object.Name}}) Inject(
			{{- range $field := $object.Fields }}
				{{- if $field.IsResolver }}
					{{lcFirst $object.Name}}{{ $field.GoFieldName}} *{{ lookupImport (gpkg $object.Name $field.Name $field) }}.{{ gtyp $object.Name $field.Name }},
				{{- end -}}
			{{- end }}
		) {
			{{- range $field := $object.Fields }}
				{{- if $field.IsResolver }}
					r.resolve{{ $field.GoFieldName}} = {{lcFirst $object.Name}}{{ $field.GoFieldName}}.{{ gmethod $object.Name $field.Name}}
				{{- end -}}
			{{- end }}
		}

		{{ range $field := $object.Fields -}}
			{{- if $field.IsResolver -}}
				func (r *{{lcFirst $root.TypeName}}{{$object.Name}}) {{$field.GoFieldName}}{{ $field.ShortResolverDeclaration }} {
					return r.resolve{{$field.GoFieldName}}(ctx,
						{{- if not $object.Root }}obj,{{end -}}
						{{- range $arg := $field.Args}}
							{{- $arg.VarName}},
						{{- end }}
					)
				}
			{{ end -}}
		{{ end -}}
	{{ end -}}
{{ end }}

func direct(root *{{$root.TypeName}}) map[string]interface{} {
	return map[string]interface{}{
	{{ range $object := .Objects -}}
		{{- if $object.HasResolvers -}}
			{{ range $field := $object.Fields -}}
				{{- if $field.IsResolver -}}
					"{{$object.Name}}.{{$field.GoFieldName}}": root.{{$object.Name}}().{{$field.GoFieldName}},
				{{ end -}}
			{{ end -}}
		{{ end -}}
	{{ end }}
	}
}
`,
	})
}

type resolverBuild struct {
	*codegen.Data

	TypeName string
}
