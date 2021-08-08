package synkr

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type CreateRequestBody struct {
	Url string
}

type RequestLynk struct {
	Id string
	Url string
}

type SynkrRouter struct {
	r *mux.Router
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

func NewRouter() *SynkrRouter {
	r := mux.NewRouter()
	return &SynkrRouter{r:r}
}

func (s *SynkrClient) SetRoutes() {
	s.router.r.HandleFunc("/", s.rootRoute)
	s.router.r.HandleFunc("/create", s.createRoute())
	s.router.r.HandleFunc("/z/{id}", s.lynkrRoute())
	s.router.r.Use(LoggingMiddleware)
}

func (s *SynkrClient) rootRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (s *SynkrClient) createRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		l := NewRequestLynk(w, r)
		lynk, lynkErr := s.db.Save(l)
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
		lynk := s.db.FindOne(id)
		fmt.Println("found -> ", lynk.Id);
		fmt.Println("redirecting -> ", lynk.GoUrl);
		http.Redirect(w, r, lynk.GoUrl, http.StatusSeeOther)
	}
}
