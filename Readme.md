# Flamingo GraphQL

This modul allows the usage of GraphQL services in your Flamingo project.

Please see `example/` on how this works.

## Setup

Add the `graphql.Module` to your projects modules.

Generate the local code by running `go run main.go graphql`.

See the `example` folder for a complete example on how to use `flamingo.me/graphql`.

## Resolver

Everything needs to be resolved, unless `gqlgen` can figure something out on its own.
Best practive here is to provide default resolvers in your code, which can be used via embedding.

See `example/user` for an example UserQuery resolver.
