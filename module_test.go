package graphql_test

import (
	"testing"

	"flamingo.me/flamingo/v3/framework/config"

	"flamingo.me/graphql"
)

func TestModule_Configure(t *testing.T) {
	t.Parallel()

	if err := config.TryModules(nil, new(graphql.Module)); err != nil {
		t.Error(err)
	}
}
