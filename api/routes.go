package api

import (
	"fmt"
	"net/http"
	"github.com/mattleong/lynkr/logger"
	"github.com/gorilla/mux"
)

func RootRoute(w http.ResponseWriter, req *http.Request) {
    w.WriteHeader(http.StatusOK)
	logger.Log("hit: /")
	fmt.Fprintf(w, "hello\n")
}

func CreateRoute(w http.ResponseWriter, req *http.Request) {
    w.WriteHeader(http.StatusOK)
	logger.Log("hit: /create")
	fmt.Fprintf(w, "create\n")
}

func LynkrRoute(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "link id: %v\n", vars["id"])
}
