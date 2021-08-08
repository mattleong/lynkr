package synkr

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// type SynkrRouter *mux.Router

type CreateRequestBody struct {
	Url string
}

type RequestLynk struct {
	Id string
	Url string
}

func NewRequestLynk(w http.ResponseWriter, r *http.Request) *RequestLynk {
	var body CreateRequestBody
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}

	id := GenerateId(10)

	return &RequestLynk{ Id: id, Url: body.Url }
}

func NewRouter() *mux.Router {
	return mux.NewRouter()
}

func (s *SynkrClient) SetRoutes() {
	s.router.HandleFunc("/", s.rootRoute)
	s.router.HandleFunc("/create", s.createRoute())
	s.router.HandleFunc("/z/{id}", s.lynkrRoute())
}

func (s *SynkrClient) rootRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (s *SynkrClient) createRoute() http.HandlerFunc {
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

func (s *SynkrClient) lynkrRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		lynk := s.FindOne(id)
		fmt.Println("found -> ", lynk.Id);
		fmt.Println("redirecting -> ", lynk.GoUrl);
		http.Redirect(w, r, lynk.GoUrl, http.StatusSeeOther)
	}
}
