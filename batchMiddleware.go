package graphql

import (
	"context"

	gql "github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
)

const (
	sameOperationsDefaultThreshold = 2
	allOperationsDefaultThreshold  = 10
)

func BatchMiddleware(
	cfg *struct {
		SameOperationsThreshold int `inject:"config:graphql.batchMiddleware.sameOperationsThreshold,optional"`
		AllOperationsThreshold  int `inject:"config:graphql.batchMiddleware.sameOperationsThreshold,optional"`
	},
) func(ctx context.Context, next gql.OperationHandler) gql.ResponseHandler {
	return func(ctx context.Context, next gql.OperationHandler) gql.ResponseHandler {
		var sameOperationsThreshold int
		var allOperationsThreshold int

		if cfg == nil {
			sameOperationsThreshold = sameOperationsDefaultThreshold
			allOperationsThreshold = allOperationsDefaultThreshold
		} else {
			sameOperationsThreshold = cfg.SameOperationsThreshold
			allOperationsThreshold = cfg.AllOperationsThreshold
		}

		req := gql.GetOperationContext(ctx)

		occurrences := countTopLevelGraphQLOperations(req.Operation.SelectionSet)

		countGraphqlFunctionsCalled(occurrences, req.Operation.SelectionSet)

		if isAboveThreshold(sameOperationsThreshold, allOperationsThreshold, occurrences) {
			return func(ctx context.Context) *gql.Response {
				return gql.ErrorResponse(ctx, "request not allowed")
			}
		}

		return next(ctx)
	}
}

func countTopLevelGraphQLOperations(definition []ast.Selection) map[string]int {
	mapOfOccurrences := make(map[string]int)

	for _, set := range definition {
		field, ok := set.(*ast.Field)
		if !ok {
			continue
		}

		if _, exists := mapOfOccurrences[field.Name]; !exists {
			mapOfOccurrences[field.Name] = 0
		}

		mapOfOccurrences[field.Name]++
	}

	return mapOfOccurrences
}

func countGraphqlFunctionsCalled(mapOfOccurrences map[string]int, definition []ast.Selection) {
	for _, set := range definition {
		field, ok := set.(*ast.Field)
		if !ok {
			continue
		}

		// counting arguments is the only way to tell if this field is a function call
		if len(field.Arguments) != 0 {
			if _, exists := mapOfOccurrences[field.Name]; !exists {
				mapOfOccurrences[field.Name] = 0
			}

			mapOfOccurrences[field.Name]++
		}

		if len(field.SelectionSet) > 0 {
			countGraphqlFunctionsCalled(mapOfOccurrences, field.SelectionSet)
		}
	}
}

func isAboveThreshold(threshold, operationsThreshold int, operations map[string]int) bool {
	if len(operations) > operationsThreshold {
		return true
	}

	for _, operationsNumber := range operations {
		if operationsNumber > threshold {
			return true
		}
	}

	return false
}
