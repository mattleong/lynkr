package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mattleong/lynkr/pkg/db"
)

type Router struct {
	r *mux.Router
}

type createRequestBody struct {
	Url string
}

func NewRouter() Router {
	r := mux.NewRouter()
	return Router{r:r}
}

func (router *Router) SetRoutes(db db.DatabaseStore) {
	router.r.HandleFunc("/", router.rootRoute)
	router.r.HandleFunc("/create", router.createRoute(db))
	router.r.HandleFunc("/z/{id}", router.lynkrRoute(db))
	router.r.Use(loggingMiddleware)
	http.Handle("/", router.r)
}

func (s *Router) rootRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (s *Router) createRoute(db db.DatabaseStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		ctx := r.Context()
		var body createRequestBody
		err := json.NewDecoder(r.Body).Decode(&body)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		lynk, err := db.SaveLynk(ctx, body.Url)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Fatal(err)
		}

		log.Printf("Created: %s -> %s\n", lynk.Id, lynk.GoUrl);
		res, _ := json.Marshal(lynk)
		w.Write(res)
	}
}

func (s *Router) lynkrRoute(db db.DatabaseStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		ctx := r.Context()
		lynk, err := db.FindLynkById(ctx, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Fatal(err)
		}

		log.Printf("Found: %s redirecting -> %s\n", lynk.Id, lynk.GoUrl);
		http.Redirect(w, r, lynk.GoUrl, http.StatusSeeOther)
	}
}
