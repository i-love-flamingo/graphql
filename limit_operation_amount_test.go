package graphql_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/99designs/gqlgen/graphql/handler/testserver"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/stretchr/testify/assert"

	"flamingo.me/graphql"
)

func Test_LimitOperationAmountMiddleware(t *testing.T) {
	t.Parallel()

	t.Run("deny when there is too many same operations called", func(t *testing.T) {
		t.Parallel()

		srv := testserver.New()

		srv.AddTransport(transport.GET{})
		srv.AddTransport(transport.POST{})

		srv.AroundOperations(graphql.LimitOperationAmountMiddleware(
			&struct {
				SameOperationLimit  int `inject:"config:graphql.security.limitOperationAmount.sameOperationLimit,optional"`
				TotalOperationLimit int `inject:"config:graphql.security.limitOperationAmount.totalOperationLimit,optional"`
			}{
				SameOperationLimit:  2,
				TotalOperationLimit: 10,
			}))

		body := `{
				  "query": "query { user1: name user2: name user3: name user4: name user5: name }"
				}`

		resp := doRequest(srv, "POST", "/graphql", body)
		assert.Equal(t, http.StatusOK, resp.Code, resp.Body.String())
		assert.Equal(t, `{"errors":[{"message":"request not allowed"}],"data":null}`, resp.Body.String())
	})

	t.Run("deny when there are too many different operations invoked in one query", func(t *testing.T) {
		t.Parallel()

		srv := testserver.New()

		srv.AddTransport(transport.GET{})
		srv.AddTransport(transport.POST{})

		srv.AroundOperations(graphql.LimitOperationAmountMiddleware(
			&struct {
				SameOperationLimit  int `inject:"config:graphql.security.limitOperationAmount.sameOperationLimit,optional"`
				TotalOperationLimit int `inject:"config:graphql.security.limitOperationAmount.totalOperationLimit,optional"`
			}{
				SameOperationLimit:  27,
				TotalOperationLimit: 0,
			}))

		body := `{
				  "query": "query { user1: name user2: name user3: name user4: name user5: name }"
				}`

		resp := doRequest(srv, "POST", "/graphql", body)
		assert.Equal(t, http.StatusOK, resp.Code, resp.Body.String())
		assert.Equal(t, `{"errors":[{"message":"request not allowed"}],"data":null}`, resp.Body.String())
	})

	t.Run("allow when request is below both thresholds", func(t *testing.T) {
		t.Parallel()

		srv := testserver.New()

		srv.AddTransport(transport.GET{})
		srv.AddTransport(transport.POST{})

		srv.AroundOperations(graphql.LimitOperationAmountMiddleware(
			&struct {
				SameOperationLimit  int `inject:"config:graphql.security.limitOperationAmount.sameOperationLimit,optional"`
				TotalOperationLimit int `inject:"config:graphql.security.limitOperationAmount.totalOperationLimit,optional"`
			}{
				SameOperationLimit:  10,
				TotalOperationLimit: 10,
			}))

		body := `{
				  "query": "query { user1: name user2: name }"
				}`

		resp := doRequest(srv, "POST", "/graphql", body)
		assert.Equal(t, http.StatusOK, resp.Code, resp.Body.String())
		assert.Equal(t, `{"data":{"name":"test"}}`, resp.Body.String())
	})
}

func doRequest(handler http.Handler, method string, target string, body string) *httptest.ResponseRecorder {
	r := httptest.NewRequest(method, target, strings.NewReader(body))

	r.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	handler.ServeHTTP(w, r)

	return w
}
