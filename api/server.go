package api

import (
	"net/http"
	"github.com/mattleong/lynkr/logger"
	"github.com/gorilla/mux"
)

func ServerStart() {
    r := mux.NewRouter()
	logger.Log("Lynkr server started...")

	r.HandleFunc("/", RootRoute)
	r.HandleFunc("/create", CreateRoute)
	r.HandleFunc("/z/{id}", LynkrRoute)
	http.Handle("/", r)
	http.ListenAndServe(":3000", nil)
}
