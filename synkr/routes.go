package synkr

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *SynkrClient) RootRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (s *SynkrClient) CreateRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		l := NewRequestLynk(w, r)
		lynk, lynkErr := s.Save(l)
		if lynkErr != nil {
			return
		}

		res, _ := json.Marshal(lynk)
		w.Write(res)
	}
}

func (s *SynkrClient) LynkrRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
}
