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
		switch v := v.(type) {
		case ModelMapEntry:
			t := reflect.TypeOf(v.Type)
			if t.Kind() == reflect.Ptr {
				t = t.Elem()
			}

			fields := make(map[string]config.TypeMapField, len(v.Fields))
			for k, v := range v.Fields {
				fields[k] = config.TypeMapField{FieldName: v}
			}

			res[k] = config.TypeMapEntry{
				Model:  []string{t.PkgPath() + "." + t.Name()},
				Fields: fields,
			}

		default:
			t := reflect.TypeOf(v)
			if t.Kind() == reflect.Ptr {
				t = t.Elem()
			}
			res[k] = config.TypeMapEntry{
				Model: []string{t.PkgPath() + "." + t.Name()},
			}
		}
	}
	return res
}

type ModelMapEntry struct {
	Type   interface{}
	Fields map[string]string
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
