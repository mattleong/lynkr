package synkr

import (
	"fmt"
	"net/http"
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
	s := NewSynkrClient()
	s.SetRoutes()
	defer s.db.disconnect()

	s.router.Use(loggingMiddleware)

	http.Handle("/", s.router)
	httpErr := http.ListenAndServe(":3000", nil)
	if httpErr != nil {
		return
	}
}
