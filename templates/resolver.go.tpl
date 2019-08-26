// +build !graphql

package graphql

import "flamingo.me/graphql"

// this file is a starting point for the project specific resolvers
// it will not be regenerated!

type rootResolver struct {
	queryResolver    *queryResolver
	mutationResolver *mutationResolver
}

// interface guard
var _ ResolverRoot = new(rootResolver)

// Inject root resolver
func (r *rootResolver) Inject(queryResolver *queryResolver, mutationResolver *mutationResolver) {
	r.queryResolver = queryResolver
	r.mutationResolver = mutationResolver
}

// Query getter
func (r *rootResolver) Query() QueryResolver {
	return r.queryResolver
}

// Mutation getter
func (r *rootResolver) Mutation() MutationResolver {
	return r.mutationResolver
}

type queryResolver struct {
	*graphql.FlamingoQueryResolver
}

// Inject dependencies
func (r *queryResolver) Inject(flamingoQueryResolver *graphql.FlamingoQueryResolver) {
	r.FlamingoQueryResolver = flamingoQueryResolver
}

type mutationResolver struct {
	*graphql.FlamingoQueryResolver
}

// Inject dependencies
func (r *mutationResolver) Inject(flamingoQueryResolver *graphql.FlamingoQueryResolver) {
	r.FlamingoQueryResolver = flamingoQueryResolver
}
