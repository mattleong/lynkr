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

func (router *LynkrRouter) setRoutes(db *LynkrDB) {
	router.r.HandleFunc("/", router.rootRoute)
	router.r.HandleFunc("/create", router.createRoute(db))
	router.r.HandleFunc("/z/{id}", router.lynkrRoute(db))
	router.r.Use(loggingMiddleware)
	http.Handle("/", router.r)
}

func (s *LynkrRouter) rootRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (s *LynkrRouter) createRoute(db *LynkrDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		ctx := r.Context()
		l := NewRequestLynk(w, r)
		lynk, err := db.SaveLynk(ctx, l)
		if err != nil {
			log.Println(err)
		}

		log.Printf("Created: %s -> %s\n", lynk.Id, lynk.GoUrl);
		res, _ := json.Marshal(lynk)
		w.Write(res)
	}
}

func (s *LynkrRouter) lynkrRoute(db *LynkrDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		ctx := r.Context()
		lynk, err := db.FindLynkById(ctx, id)
		if err != nil {
			log.Println(err)
		}

		log.Printf("Found: %s redirecting -> %s\n", lynk.Id, lynk.GoUrl);
		http.Redirect(w, r, lynk.GoUrl, http.StatusSeeOther)
	}
}
