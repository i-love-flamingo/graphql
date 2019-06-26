package graphql

import (
	"context"
	"reflect"

	"github.com/99designs/gqlgen/codegen/config"
)

// ModelMap is a helper to quickly create map[string]TypeMapEntry for types
type ModelMap map[string]interface{}

// Models creates the references via reflection
func (m ModelMap) Models() map[string]config.TypeMapEntry {
	res := make(map[string]config.TypeMapEntry, len(m))
	for k, v := range m {
		v := reflect.TypeOf(v)
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		res[k] = config.TypeMapEntry{
			Model: []string{v.PkgPath() + "." + v.Name()},
		}
	}
	return res
}

type FlamingoQueryResolver struct{}

func (*FlamingoQueryResolver) Flamingo(ctx context.Context) (*string, error) {
	return FlamingoResolver(ctx)
}

// FlamingoResolver example
func FlamingoResolver(_ context.Context) (*string, error) {
	flamingo := "flamingo"
	return &flamingo, nil
}
