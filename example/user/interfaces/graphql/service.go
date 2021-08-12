package graphql

import (
	"flamingo.me/graphql"
	"flamingo.me/graphql/example/user/domain"
)

//go:generate go run github.com/go-bindata/go-bindata/v3/go-bindata -nometadata -o fs.go -pkg graphql schema.graphql

// Service for graphql
type Service struct{}

// Schema defines the graphql schema
func (*Service) Schema() []byte {
	return MustAsset("schema.graphql")
}

// Types set up the GraphQL to Go Type mappings
func (*Service) Types(types *graphql.Types) {
	types.Map("User", domain.User{})
	types.Map("User_Attributes", domain.Attributes{})
	types.Resolve("Query", "User", UserQueryResolver{}, "User")
}
