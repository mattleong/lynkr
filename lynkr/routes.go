package lynkr

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	r *mux.Router
}

func newRouter() *Router {
	r := mux.NewRouter()
	return &Router{r:r}
}

func (router *Router) setRoutes(db *Database) {
	router.r.HandleFunc("/", router.rootRoute)
	router.r.HandleFunc("/create", router.createRoute(db))
	router.r.HandleFunc("/z/{id}", router.lynkrRoute(db))
	router.r.Use(loggingMiddleware)
	http.Handle("/", router.r)
}

func (s *Router) rootRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (s *Router) createRoute(db *Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		ctx := r.Context()
		l := NewRequestLynk(w, r)
		lynk, err := db.SaveLynk(ctx, l)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Created: %s -> %s\n", lynk.Id, lynk.GoUrl);
		res, _ := json.Marshal(lynk)
		w.Write(res)
	}
}

func (s *Router) lynkrRoute(db *Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		ctx := r.Context()
		lynk, err := db.FindLynkById(ctx, id)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Found: %s redirecting -> %s\n", lynk.Id, lynk.GoUrl);
		http.Redirect(w, r, lynk.GoUrl, http.StatusSeeOther)
	}
}
