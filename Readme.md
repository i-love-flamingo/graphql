# Flamingo GraphQL

This modul allows the usage of GraphQL services in your Flamingo project.

Please see `example/` on how this works.

## Setup

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

## Resolver

Everything needs to be resolved, unless `gqlgen` can figure something out on its own.
Best practive here is to provide default resolvers in your code, which can be used via embedding.

See `example/user` for an example UserQuery resolver.
