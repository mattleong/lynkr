package routes

import (
	"log"
	"net/http"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Hit: %s\n", r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
