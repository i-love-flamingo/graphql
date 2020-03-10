package graphql

import (
	"flamingo.me/graphql"
	"flamingo.me/graphql/example/user/domain"
	"github.com/99designs/gqlgen/codegen/config"
)

//go:generate go run github.com/go-bindata/go-bindata/go-bindata -nometadata -o fs.go -pkg graphql schema.graphql

// Service service for graphql
type Service struct{}

// Schema defines the graphql schema
func (*Service) Schema() []byte {
	return MustAsset("schema.graphql")
}

// Models map associations between graphql and go
func (*Service) Models() map[string]config.TypeMapEntry {
	return graphql.ModelMap{
		"User": domain.User{},
	}.Models()
}
