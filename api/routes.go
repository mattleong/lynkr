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

func CreateRoute(s *synkr.SynkrClient) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		l := lynkr.NewRequestLynk(w, r)
		lynk, lynkErr := s.Save(l)
		if lynkErr != nil {
			return
		}

		res, _ := json.Marshal(lynk)
		w.Write(res)
	}

	return fn
}

func LynkrRoute(s *synkr.SynkrClient) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)
		id := vars["id"]
		lynk := s.FindOne(id)
		fmt.Println("return lynk!!: ", lynk.Url);
		//	http.Redirect(w, r, "https://google.com", http.StatusSeeOther)
		res, _ := json.Marshal(lynk)
		w.Write(res)
	}

	return fn
}
