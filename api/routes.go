package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mattleong/lynkr/lynkr"
	"github.com/mattleong/lynkr/synkr"
	"net/http"
)

func RootRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func CreateRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	l := lynkr.NewLynkFromRequest(w, r)
	lynk, lynkErr := synkr.SaveLynk(l)
	if lynkErr != nil {
		return
	}
	res, _ := json.Marshal(lynk)
	w.Write(res)
}

func LynkrRoute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintf(w, "hit: /z/%s\n", id)
	//	http.Redirect(w, r, "https://google.com", http.StatusSeeOther)
}
