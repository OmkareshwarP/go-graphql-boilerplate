package utils

import (
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//handle OPTIONS requests for CORS
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Headers", "*")
			w.Header().Set("Access-Control-Allow-Methods", "*")
			w.WriteHeader(http.StatusNoContent)
			return
		}
		if r.URL.Path == "/health" {
			// Skip authentication for health
			next.ServeHTTP(w, r)
			return
		}
		if r.Method == "GET" {
			// Skip authentication for introspection/playground
			next.ServeHTTP(w, r)
			return
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}
