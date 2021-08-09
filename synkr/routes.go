package synkr

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type SynkrRouter struct {
	r *mux.Router
}

type RequestLynk struct {
	Id string
	Url string
}

type createRequestBody struct {
	Url string
}

func newRequestLynk(w http.ResponseWriter, r *http.Request) *RequestLynk {
	var body createRequestBody
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}

	id := GenerateId(10)

	return &RequestLynk{ Id: id, Url: body.Url }
}

func newRouter() *SynkrRouter {
	r := mux.NewRouter()
	return &SynkrRouter{r:r}
}

func (s *SynkrClient) setRoutes() {
	s.router.r.HandleFunc("/", s.rootRoute)
	s.router.r.HandleFunc("/create", s.createRoute())
	s.router.r.HandleFunc("/z/{id}", s.lynkrRoute())
	s.router.r.Use(loggingMiddleware)
}

func (s *SynkrClient) rootRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (s *SynkrClient) createRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		ctx := r.Context()
		l := newRequestLynk(w, r)
		lynk, lynkErr := s.SaveLynk(ctx, l)
		if lynkErr != nil {
			log.Fatal(lynkErr)
		}

		fmt.Printf("created: %s -> %s\n", lynk.Id, lynk.GoUrl);
		res, _ := json.Marshal(lynk)
		w.Write(res)
	}
}

func (s *SynkrClient) lynkrRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		ctx := r.Context()
		lynk, lynkErr := s.FindLynkById(ctx, id)
		if lynkErr != nil {
			log.Fatal(lynkErr)
		}

		fmt.Printf("found: %s redirecting -> %s\n", lynk.Id, lynk.GoUrl);
		http.Redirect(w, r, lynk.GoUrl, http.StatusSeeOther)
	}
}
