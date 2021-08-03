package api

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/mattleong/lynkr/lynkr"
	"encoding/json"
)

func RootRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func CreateRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	lynk := lynkr.NewLynkFromRequest(w, r)
	res, _ := json.Marshal(lynk)
	w.Write(res)
}

func LynkrRoute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintf(w, "hit: /z/%s\n", id)
	//	http.Redirect(w, r, "https://google.com", http.StatusSeeOther)
}
