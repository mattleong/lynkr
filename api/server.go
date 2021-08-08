package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mattleong/lynkr/synkr"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Hit: %s\n", r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func ServerStart() {
	fmt.Println("Lynkr server started...")

	// set up synkr client
	s := synkr.NewSynkrClient()
	s.Ping()

	r := mux.NewRouter()

	r.HandleFunc("/", s.RootRoute)
	// pass db to as route dep
	r.HandleFunc("/create", s.CreateRoute())
	r.HandleFunc("/z/{id}", s.LynkrRoute())
	r.Use(loggingMiddleware)

	http.Handle("/", r)
	httpErr := http.ListenAndServe(":3000", nil)
	if httpErr != nil {
		return
	}
}
