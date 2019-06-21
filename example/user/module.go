package user

import (
	"flamingo.me/dingo"
	"flamingo.me/graphql"
	"flamingo.me/graphql/example/user/domain"
	"flamingo.me/graphql/example/user/infrastructure"
	graphqlinterface "flamingo.me/graphql/example/user/interfaces/graphql"
)

type Module struct{}

func (*Module) Configure(injector *dingo.Injector) {
	injector.BindMulti(new(graphql.Service)).To(new(graphqlinterface.Service))
	injector.Bind(new(domain.UserService)).To(infrastructure.UserServiceImpl{})
}

func (*Module) Depends() []dingo.Module {
	return []dingo.Module{
		new(graphql.Module),
	}
}
