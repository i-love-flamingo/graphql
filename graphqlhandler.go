package graphql

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"

	"flamingo.me/flamingo/v3/framework/web"
)

type (
	gqlHandler struct {
		recorder *httptest.ResponseRecorder
		request  *web.Request
	}
)

// wrapGqlHandler wraps an http.Handler to be used in the flamingo http package
func wrapGqlHandler(handler http.Handler) web.Action {
	return func(ctx context.Context, req *web.Request) web.Result {
		rw := httptest.NewRecorder()
		handler.ServeHTTP(rw, req.Request().WithContext(ctx))

		return &gqlHandler{
			request:  req,
			recorder: rw,
		}
	}
}

func (h *gqlHandler) Apply(_ context.Context, rw http.ResponseWriter) error {
	for k, vs := range h.recorder.Header() {
		for _, v := range vs {
			rw.Header().Add(k, v)
		}
	}

	_, err := io.Copy(rw, h.recorder.Body)
	if err != nil {
		return fmt.Errorf("failed to write body: %w", err)
	}

	return nil
}
