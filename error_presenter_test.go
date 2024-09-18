package graphql_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vektah/gqlparser/v2/gqlerror"

	"flamingo.me/graphql"
)

func Test_DropTypeHintsFromErrorMessage(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx context.Context
		err error
	}

	tests := []struct {
		name string
		args args
		want *gqlerror.Error
	}{
		{
			name: "error message for validation rule without specific presenter is returned as is",
			args: args{
				err: &gqlerror.Error{
					Message: "some message",
					Rule:    "NoUnusedVariables",
				},
			},
			want: &gqlerror.Error{
				Message: "some message",
				Rule:    "NoUnusedVariables",
			},
		},
		{
			name: `error message for validation rule "FieldsOnCorrectType" is trimmed to prevent exposing correct types`,
			args: args{
				err: &gqlerror.Error{
					Message: `Cannot query field "Commerce_Cart" on type "Query". Did you mean "Commerce_Product" or "Commerce_Customer"?`,
					Rule:    "FieldsOnCorrectType",
				},
			},
			want: &gqlerror.Error{
				Message: `Cannot query field "Commerce_Cart" on type "Query".`,
				Rule:    "FieldsOnCorrectType",
			},
		},
		{
			name: `Non-GraphQL error is wrapped as GraphQL error`,
			args: args{
				err: fmt.Errorf("some random error"),
			},
			want: &gqlerror.Error{
				Err:     fmt.Errorf("some random error"),
				Message: `some random error`,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equalf(t, tt.want, graphql.DropTypeHintsFromErrorMessage(tt.args.ctx, tt.args.err), "DropTypeHintsFromErrorMessage(%v, %v)", tt.args.ctx, tt.args.err)
		})
	}
}
