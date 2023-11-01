package graphql

import (
	"net/http"
)

type corsHandler struct {
	origins []string
}

func (h *corsHandler) validateOrigin(origin string) bool {
	for _, allowedOrigin := range h.origins {
		if allowedOrigin == origin || allowedOrigin == "*" {
			return true
		}
	}

	return false
}

func (h *corsHandler) addCorsHeaders(rw http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")
	if h.validateOrigin(origin) {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Credentials", "true")
		rw.Header().Set("Access-Control-Allow-Methods", "POST")
		rw.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	}
}

func (h *corsHandler) preflightHandler() http.Handler {
	return http.HandlerFunc(h.addCorsHeaders)
}

func (h *corsHandler) gqlMiddleware(gqlHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, r *http.Request) {
		h.addCorsHeaders(writer, r)
		gqlHandler.ServeHTTP(writer, r)
	})
}
