//+build !graphql

package graphql

import (
	"flamingo.me/graphql"
	"flamingo.me/graphql/example/todo"
	usergraphql "flamingo.me/graphql/example/user/interfaces/graphql"
)

// this file is a starting point for the project specific resolvers
// it will not be regenerated!

type rootResolver struct {
	todoUserResolver *todo.TodoUserResolver
	queryResolver    *queryResolver
	todoResolver     *todoResolver
}

func (r *rootResolver) Inject(
	todoUserResolver *todo.TodoUserResolver,
	queryResolver *queryResolver,
	todoResolver *todoResolver,
) *rootResolver {
	r.todoUserResolver = todoUserResolver
	r.queryResolver = queryResolver
	r.todoResolver = todoResolver
	return r
}

var _ ResolverRoot = new(rootResolver)

func (r *rootResolver) Query() QueryResolver {
	return r.queryResolver
}

func (r *rootResolver) Mutation() MutationResolver {
	panic("no mutations available")
}

func (r *rootResolver) User() UserResolver {
	return r.todoUserResolver
}

func (r *rootResolver) Todo() TodoResolver {
	return r.todoResolver
}

type queryResolver struct {
	*graphql.FlamingoQueryResolver
	*usergraphql.UserQueryResolver
}

func (r *queryResolver) Inject(flamingoQueryResolver *graphql.FlamingoQueryResolver, userQueryResolver *usergraphql.UserQueryResolver) *queryResolver {
	r.FlamingoQueryResolver = flamingoQueryResolver
	r.UserQueryResolver = userQueryResolver
	return r
}

type todoResolver struct {
	*todo.TodoResolver
}

func (r *todoResolver) Inject(tr *todo.TodoResolver) {
	r.TodoResolver = tr
}
