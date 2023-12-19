package graphql

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/99designs/gqlgen/graphql/handler/testserver"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/stretchr/testify/assert"
)

func Test_LimitQueryAmountMiddleware(t *testing.T) {
	t.Run("deny when there is too many same operations called", func(t *testing.T) {
		srv := testserver.New()

		srv.AddTransport(transport.GET{})
		srv.AddTransport(transport.POST{})

		srv.AroundOperations(LimitQueryAmountMiddleware(
			&struct {
				SameOperationsThreshold int `inject:"config:graphql.limitQueryAmountMiddleware.sameOperationsThreshold,optional"`
				AllOperationsThreshold  int `inject:"config:graphql.limitQueryAmountMiddleware.allOperationsThreshold,optional"`
			}{
				SameOperationsThreshold: 2,
				AllOperationsThreshold:  10,
			}))

		body := `{
				  "query": "query { user1: name user2: name user3: name user4: name user5: name }"
				}`

		resp := doRequest(srv, "POST", "/graphql", body)
		assert.Equal(t, http.StatusOK, resp.Code, resp.Body.String())
		assert.Equal(t, `{"errors":[{"message":"request not allowed"}],"data":null}`, resp.Body.String())
	})

	t.Run("deny when there are too many different operations invoked in one query", func(t *testing.T) {
		srv := testserver.New()

		srv.AddTransport(transport.GET{})
		srv.AddTransport(transport.POST{})

		srv.AroundOperations(LimitQueryAmountMiddleware(
			&struct {
				SameOperationsThreshold int `inject:"config:graphql.limitQueryAmountMiddleware.sameOperationsThreshold,optional"`
				AllOperationsThreshold  int `inject:"config:graphql.limitQueryAmountMiddleware.allOperationsThreshold,optional"`
			}{
				SameOperationsThreshold: 27,
				AllOperationsThreshold:  0,
			}))

		body := `{
				  "query": "query { user1: name user2: name user3: name user4: name user5: name }"
				}`

		resp := doRequest(srv, "POST", "/graphql", body)
		assert.Equal(t, http.StatusOK, resp.Code, resp.Body.String())
		assert.Equal(t, `{"errors":[{"message":"request not allowed"}],"data":null}`, resp.Body.String())
	})

	t.Run("allow when request is below both thresholds", func(t *testing.T) {
		srv := testserver.New()

		srv.AddTransport(transport.GET{})
		srv.AddTransport(transport.POST{})

		srv.AroundOperations(LimitQueryAmountMiddleware(
			&struct {
				SameOperationsThreshold int `inject:"config:graphql.limitQueryAmountMiddleware.sameOperationsThreshold,optional"`
				AllOperationsThreshold  int `inject:"config:graphql.limitQueryAmountMiddleware.allOperationsThreshold,optional"`
			}{
				SameOperationsThreshold: 10,
				AllOperationsThreshold:  10,
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
