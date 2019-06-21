package graphql

import (
	"flamingo.me/graphql"
	"flamingo.me/graphql/example/user/domain"
	"github.com/99designs/gqlgen/codegen/config"
)

//go:generate go run github.com/go-bindata/go-bindata/go-bindata -o fs.go -pkg graphql schema.graphql

type Service struct{}

func (*Service) Schema() []byte {
	return MustAsset("schema.graphql")
}

func (*Service) Models() map[string]config.TypeMapEntry {
	return graphql.ModelMap{
		"User": domain.User{},
	}.Models()
}
