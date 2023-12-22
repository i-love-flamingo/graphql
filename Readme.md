# Flamingo GraphQL

This module allows the usage of GraphQL services in your Flamingo project.

## Example
The module contains a working example that you can try out yourself. 
See https://github.com/i-love-flamingo/graphql/tree/master/example

## Use Graphql in your Flaming project

Create a new Module and bind a `graphql.Service`.

The service has both the Schema and Model-Mapping information.

It also makes sense to declare `graphql.Module` as a direct dependency from your model:
```go
package todo

import (
	"flamingo.me/dingo"
	"flamingo.me/graphql"
	"flamingo.me/graphql/example/todo/domain"
	"flamingo.me/graphql/example/user"
	"github.com/99designs/gqlgen/codegen/config"
)

// Service for the todo graphql service
type service struct{}

// Schema getter for graphql
func (*service) Schema() []byte {
	// language=graphql
	return []byte(`
type Todo {
	id: ID!
	task: String!
	done: Boolean!
}
extend type User {
	todos: [Todo]
}
extend type Mutation {
	TodoAdd(user: ID!, task: String!): Todo
	TodoDone(todo: ID!, done: Boolean!): Todo
}
`)
}

// Types mapping between graphql and go
func (*service) Types(types *graphql.Types) {
	types.Map("Todo", domain.Todo{})
	config.Resolve("User", "todos", UserResolver{}, "Todos")
	config.Resolve("Mutation", "TodoAdd", MutationResolver{}, "TodoAdd")
	config.Resolve("Mutation", "TodoDone", MutationResolver{}, "TodoDone")
}

// Module struct for the todo module
type Module struct{}

// Configure registers the service graphql service
func (Module) Configure(injector *dingo.Injector) {
	injector.BindMulti(new(graphql.Service)).To(new(service))
}

// Depends marks dependency to the graphql Module
func (*Module) Depends() []dingo.Module {
	return []dingo.Module{
		new(graphql.Module),
	}
}
```

## Setup Graphql in your Flamingo project

Add the `graphql.Module` to your projects modules.

Generate the local code by running `go run main.go graphql`.
Better is to add the following lines to your project main file:

```go
//go:generate rm -f graphql/generated.go
//go:generate go run -tags graphql main.go graphql
```

Then you can run `go generate .` instead of `go run main.go graphql`

After this the project specific graphql code is generated inside the folder `graphql`

You then also need to add the generated local Module to your main project file:

```

import (
    projectGraphql "yourpoject-namespace/graphql"
)

....
   new(projectGraphql.Module),
....

```
Now when you start your application you are able to use the graphql API and the graphql console.
See the `example` folder for a complete example on how to use `flamingo.me/graphql`.

The next thing you need to do is to add Resolver to the file `resolver.go`:

### Resolver and Modifier

Everything needs to be resolved, unless `gqlgen` can figure something out on its own.

Best practice here is to provide default resolvers in your code, which can be used via embedding.
See [`example/user`](https://github.com/i-love-flamingo/graphql/blob/master/example/user/interfaces/graphql/resolver.go) for an example UserQuery resolver.

### Suggested Naming Conventions

We recommend to namespace your Types and Type extensions with the Project name. 
For Flamingo Core Framework GraphQL Schema we use the prefix `Core_` and for Flamingo Commerce we use `Commerce_`
 

## Config

You can enable `LimitOperationAmountMiddleware` to prevent batching attack by setting `graphql.security.limitOperationAmount.enable` to true. 

`graphql.security.limitOperationAmount.sameOperationLimit` option can be used to set a limit for the same operations called in a single request.

`graphql.security.limitOperationAmount.totalOperationLimit` option can be used to set a limit for all the operations called in a single request.



## Resources

Learn GraphQL: https://graphql.org/learn/

GraphQL Specification: https://graphql.github.io/graphql-spec/June2018/#sec-Schema
