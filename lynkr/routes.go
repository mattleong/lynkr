package lynkr

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type LynkrRouter struct {
	r *mux.Router
}

func newRouter() *LynkrRouter {
	r := mux.NewRouter()
	return &LynkrRouter{r:r}
}

func (s *LynkrClient) setRoutes() {
	s.router.r.HandleFunc("/", s.rootRoute)
	s.router.r.HandleFunc("/create", s.createRoute())
	s.router.r.HandleFunc("/z/{id}", s.lynkrRoute())
	s.router.r.Use(loggingMiddleware)
}

func (s *LynkrClient) rootRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (s *LynkrClient) createRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		ctx := r.Context()
		l := NewRequestLynk(w, r)
		lynk, lynkErr := s.SaveLynk(ctx, l)
		if lynkErr != nil {
			log.Fatal(lynkErr)
		}

		fmt.Printf("created: %s -> %s\n", lynk.Id, lynk.GoUrl);
		res, _ := json.Marshal(lynk)
		w.Write(res)
	}
}

func (s *LynkrClient) lynkrRoute() http.HandlerFunc {
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
