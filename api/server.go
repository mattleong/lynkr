package api

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("hit: %s\n", r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func ServerStart() {
	fmt.Println("Lynkr server started...")

	r := mux.NewRouter()

	r.HandleFunc("/", RootRoute)
	r.HandleFunc("/create", CreateRoute)
	r.HandleFunc("/z/{id}", LynkrRoute)
	r.Use(loggingMiddleware)

	http.Handle("/", r)
	http.ListenAndServe(":3000", nil)
}
