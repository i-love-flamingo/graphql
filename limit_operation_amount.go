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

var _ gql.OperationMiddleware = LimitOperationAmountMiddleware(nil)

func LimitOperationAmountMiddleware(
	cfg *struct {
		SameOperationLimit  int `inject:"config:graphql.security.limitOperationAmount.sameOperationLimit,optional"`
		TotalOperationLimit int `inject:"config:graphql.security.limitOperationAmount.totalOperationLimit,optional"`
	},
) func(ctx context.Context, next gql.OperationHandler) gql.ResponseHandler {
	return func(ctx context.Context, next gql.OperationHandler) gql.ResponseHandler {
		sameOperationLimit := sameOperationsDefaultThreshold
		totalOperationLimit := allOperationsDefaultThreshold

		if cfg != nil {
			sameOperationLimit = cfg.SameOperationLimit
			totalOperationLimit = cfg.TotalOperationLimit
		}

		req := gql.GetOperationContext(ctx)

		occurrences := countTopLevelGraphQLOperations(req.Operation.SelectionSet)

		if isAboveThreshold(sameOperationLimit, totalOperationLimit, occurrences) {
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

func isAboveThreshold(sameOperationLimit, totalOperationLimit int, operations map[string]int) bool {
	if len(operations) > totalOperationLimit {
		return true
	}

	for _, operationsNumber := range operations {
		if operationsNumber > sameOperationLimit {
			return true
		}
	}

	return false
}
