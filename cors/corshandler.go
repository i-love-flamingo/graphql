package cors

import (
	"net/http"
)

type Handler struct {
	Whitelist []string
}

func (h *Handler) validateOrigin(origin string) bool {
	for _, whitelistEntry := range h.Whitelist {
		if whitelistEntry == origin {
			return true
		}
	}
	return false
}

func (h *Handler) addCorsHeaders(writer http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")
	if h.validateOrigin(origin) {
		writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		writer.Header().Set("Access-Control-Allow-Credentials", "true")
		writer.Header().Set("Access-Control-Allow-Methods", "POST")
		writer.Header().Set("Access-Control-Allow-Headers", "content-type")
	}
}

func (h *Handler) PreflightHandler() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, r *http.Request) {
		h.addCorsHeaders(writer, r)
	})
}

func (h *Handler) GqlMiddleware(gqlHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, r *http.Request) {
		h.addCorsHeaders(writer, r)
		gqlHandler.ServeHTTP(writer, r)
	})
}
