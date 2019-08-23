package graphql_test

import (
	"testing"

	"flamingo.me/dingo"
	"flamingo.me/graphql"
)

func TestModule_Configure(t *testing.T) {
	r := dingo.TryModule(new(graphql.Module))
	if r != nil {
		t.Error(r)
	}
}
