package lynkr

import (
	"encoding/json"
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
		lynk, err := s.db.SaveLynk(ctx, l)
		if err != nil {
			log.Println(err)
		}

		log.Printf("Created: %s -> %s\n", lynk.Id, lynk.GoUrl);
		res, _ := json.Marshal(lynk)
		w.Write(res)
	}
}

func (s *LynkrClient) lynkrRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		ctx := r.Context()
		lynk, err := s.db.FindLynkById(ctx, id)
		if err != nil {
			log.Println(err)
		}

		log.Printf("Found: %s redirecting -> %s\n", lynk.Id, lynk.GoUrl);
		http.Redirect(w, r, lynk.GoUrl, http.StatusSeeOther)
	}
}
