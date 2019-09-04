# Flamingo GraphQL

This modul allows the usage of GraphQL services in your Flamingo project.

## Example
The module contains a working example that you can try out yourself. 
See https://github.com/i-love-flamingo/graphql/tree/master/example

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

Best practive here is to provide default resolvers in your code, which can be used via embedding.
See [`example/user`](https://github.com/i-love-flamingo/graphql/blob/master/example/user/interfaces/graphql/resolver.go) for an example UserQuery resolver.

### Suggested Naming Conventions

We recommend to namespace your Types and Type extensions with the Project name. 
For Flamingo Core Framework GraphQL Schema we use the prefix `Core_` and for Flamingo Commerce we use `Commerce_`
 

## Resources

Learn GraphQL: https://graphql.org/learn/

GraphQL Specification: https://graphql.github.io/graphql-spec/June2018/#sec-Schema