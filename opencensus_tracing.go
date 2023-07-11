package graphql

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"go.opencensus.io/trace"
)

var _ graphql.HandlerExtension = &OpenCensusTracingExtension{}
var _ graphql.ResponseInterceptor = &OpenCensusTracingExtension{}
var _ graphql.FieldInterceptor = &OpenCensusTracingExtension{}
var _ graphql.OperationInterceptor = &OpenCensusTracingExtension{}

type OpenCensusTracingExtension struct{}

func (o *OpenCensusTracingExtension) ExtensionName() string {
	return "OpenCensusTracing"
}

func (o *OpenCensusTracingExtension) InterceptField(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
	fieldContext := graphql.GetFieldContext(ctx)

	// for now, we only trace resolvers
	if !fieldContext.IsResolver {
		return next(ctx)
	}

	spanName := fmt.Sprintf("graphql/resolver/%s", fieldContext.Field.Name)
	if fieldContext.Object != "Query" && fieldContext.Object != "Mutation" {
		// enrich information if the resolver is registered on a graphql subtype, e.g. resolver that resolves `sortedList` on type `Users_Result`
		spanName = fmt.Sprintf("graphql/resolver/%s/%s", fieldContext.Object, fieldContext.Field.Name)
	}

	ctx, span := trace.StartSpan(ctx, spanName)
	defer span.End()

	return next(ctx)
}

func (o *OpenCensusTracingExtension) InterceptOperation(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
	return next(ctx)
}

func (o *OpenCensusTracingExtension) InterceptResponse(ctx context.Context, next graphql.ResponseHandler) *graphql.Response {
	return next(ctx)
}

func (o *OpenCensusTracingExtension) Validate(_ graphql.ExecutableSchema) error {
	return nil
}
