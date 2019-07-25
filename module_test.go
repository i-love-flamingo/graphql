package graphql_test

import (
	"flamingo.me/dingo"
	"flamingo.me/graphql"
	"testing"
)

func TestModule_Configure(t *testing.T) {
	r := dingo.TryModule(new(graphql.Module))
	if r != nil {
		t.Error(r)
	}
}
