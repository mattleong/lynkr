package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mattleong/lynkr/synkr"
	"net/http"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("hit: %s\n", r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func ServerStart() {
	fmt.Println("Lynkr server started...")

	db := synkr.GetClient()
	s := synkr.NewSynkrClient(db)

	r := mux.NewRouter()

	r.HandleFunc("/", RootRoute)
	r.HandleFunc("/create", CreateRoute(s))
	r.HandleFunc("/z/{id}", LynkrRoute)
	r.Use(loggingMiddleware)

	http.Handle("/", r)
	httpErr := http.ListenAndServe(":3000", nil)
	if httpErr != nil {
		return
	}
}
