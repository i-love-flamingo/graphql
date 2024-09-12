package graphql

import (
	"context"
	"errors"
	"strings"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

func errorPresenterFunc(_ context.Context, err error) *gqlerror.Error {
	var gqlError *gqlerror.Error

	if errors.As(err, &gqlError) {
		switch gqlError.Rule {
		case "FieldsOnCorrectType":
			// Error message of this type exposes actual GraphQL types via suggestions: e.g.: Cannot query field "Commerce_Cart" on type "Query". Did you mean "Commerce_Product" or "Commerce_Customer"?
			// We are dropping all sentences after first full stop sign (dot)
			parts := strings.SplitAfter(gqlError.Message, ".")
			if len(parts) > 0 {
				gqlError.Message = parts[0]
			}
		}

		return gqlError
	}

	return gqlerror.Wrap(err)
}
