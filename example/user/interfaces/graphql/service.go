package graphql

import (
	// embed schema.grapqhl
	_ "embed"

	"flamingo.me/graphql"
	"flamingo.me/graphql/example/user/domain"
)

//go:embed schema.graphql
var schema []byte

// Service for graphql
type Service struct{}

// Schema defines the graphql schema
func (*Service) Schema() []byte {
	return schema
}

// Types set up the GraphQL to Go Type mappings
func (*Service) Types(types *graphql.Types) {
	types.Map("User", domain.User{})
	types.Map("User_Attributes", domain.Attributes{})
	types.Resolve("Query", "User", UserQueryResolver{}, "User")
}
