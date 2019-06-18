module flamingo.me/graphql

go 1.12

require (
	flamingo.me/dingo v0.1.5
	flamingo.me/flamingo/v3 v3.0.0-beta.2.0.20190617090914-c1bec67f7a95 // indirect
	github.com/99designs/gqlgen v0.9.0
	github.com/garyburd/redigo v1.6.0 // indirect
	github.com/spf13/cobra v0.0.3
	github.com/vektah/gqlparser v1.1.2
)

replace flamingo.me/flamingo/v3 => ../flamingo/flamingo
