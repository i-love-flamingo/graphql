# Flamingo GraphQL Example

Make sure you have the graphql module with the example cloned:
```
git clone https://github.com/i-love-flamingo/graphql.git
cd graphql/example
```

Regenerate GraphQL: `go generate .`

Run Server: `go run main.go serve`

Show console: http://localhost:3322/graphql-console

Show frontend: http://localhost:3322/

## Resources

Learn GraphQL: https://graphql.org/learn/

GraphQL Specification: https://graphql.github.io/graphql-spec/June2018/#sec-Schema

## Suggested Naming Conventions

`Core_User`
`Commerce_Product`
`Project_Whatever`

## Adding new Fields/Structs

(Rerun) code generation: `go generate .`
