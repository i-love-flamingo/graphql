//+build !graphql

package graphql

// this file is a starting point for the project specific resolvers
// it will not be regenerated!

type rootResolver  struct{}

var _ ResolverRoot = new(rootResolver)

func (*rootResolver) Query() QueryResolver {
return nil
}

func (*rootResolver) Mutation() MutationResolver {
return nil
}

type queryResolver struct{}
type mutationResolver struct{}
