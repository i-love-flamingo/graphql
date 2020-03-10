package graphql

import (
	"flamingo.me/graphql"
	"flamingo.me/graphql/example/user/domain"
)

//go:generate go run github.com/go-bindata/go-bindata/v3/go-bindata -nometadata -o fs.go -pkg graphql schema.graphql

// Service service for graphql
type Service struct{}

// Schema defines the graphql schema
func (*Service) Schema() []byte {
	return MustAsset("schema.graphql")
}

// Types set up the GraphQL to Go Type mappings
func (*Service) Types(config *graphql.Types) {
	config.Map("User", domain.User{})
	config.Map("User_Attributes", domain.Attributes{})
	config.Resolve("Query", "User", UserQueryResolver{}, "User")
}
