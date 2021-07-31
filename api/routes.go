package api

import (
	"fmt"
	"net/http"
	"github.com/mattleong/lynkr/logger"
	"github.com/gorilla/mux"
)

func RootRoute(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
	logger.Log("hit: /")
	fmt.Fprintf(w, "hello\n")
}

func CreateRoute(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
	logger.Log("hit: /create")
	fmt.Fprintf(w, "create\n")
}

func LynkrRoute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Printf("hit: /z/%v\n", id)
	http.Redirect(w, r, "https://google.com", http.StatusSeeOther)
}
