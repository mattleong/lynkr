package api

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func RootRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "hello\n")
}

func CreateRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "create\n")
}

func LynkrRoute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintf(w, "hit: /z/%s\n", id)
	//	http.Redirect(w, r, "https://google.com", http.StatusSeeOther)
}
