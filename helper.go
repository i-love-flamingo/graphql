package graphql

import (
	"context"
	"reflect"
)

// Types represent information on Object->Go type mappings and resolvers
type Types struct {
	names      map[string]string
	resolver   map[string]map[string][3]string
	fields     map[string]map[string]string
	directives map[string][3]string
}

// Map references a graphql type to a go type
func (tc *Types) Map(graphqlType string, goType interface{}) {
	t := reflect.TypeOf(goType)
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if tc.names == nil {
		tc.names = make(map[string]string)
	}

	tc.names[graphqlType] = t.PkgPath() + "." + t.Name()
}

// Resolve sets a resolver on a graphqlType's graphqlField to a method of a go type
func (tc *Types) Resolve(graphqlType, graphqlField string, typ interface{}, method string) {
	if tc.resolver == nil {
		tc.resolver = make(map[string]map[string][3]string)
	}

	if tc.resolver[graphqlType] == nil {
		tc.resolver[graphqlType] = make(map[string][3]string)
	}

	t := reflect.TypeOf(typ)
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	tc.resolver[graphqlType][graphqlField] = [3]string{t.PkgPath(), t.Name(), method}
}

// GoField sets a map to find the correct go field of an object
func (tc *Types) GoField(graphqlType, graphqlField, goField string) {
	if tc.fields == nil {
		tc.fields = make(map[string]map[string]string)
	}

	if tc.fields[graphqlType] == nil {
		tc.fields[graphqlType] = make(map[string]string)
	}

	tc.fields[graphqlType][graphqlField] = goField
}

// Directive specifies a directive resolver for a graphql directive
func (tc *Types) Directive(graphqlDirective string, typ interface{}, method string) {
	if tc.directives == nil {
		tc.directives = make(map[string][3]string)
	}

	t := reflect.TypeOf(typ)
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	tc.directives[graphqlDirective] = [3]string{t.PkgPath(), t.Name(), method}
}

// FlamingoQueryResolver always resolves to the string "flamingo" for the default schemas.
type FlamingoQueryResolver struct{}

// Flamingo field resolver
func (*FlamingoQueryResolver) Flamingo(ctx context.Context) (*string, error) {
	return FlamingoResolver(ctx)
}

// FlamingoResolver always returns "flamingo" for default schemas.
func FlamingoResolver(_ context.Context) (*string, error) {
	flamingo := "flamingo"
	return &flamingo, nil
}
