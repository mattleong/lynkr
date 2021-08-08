package synkr

import (
	"context"
	"time"

	"github.com/gorilla/mux"
)

type SynkrClient struct {
	db *SynkrDB
	router *mux.Router
}

func CreateContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	return ctx, cancel
}

func NewSynkrClient() *SynkrClient {
	db := NewDBClient()
	r := NewRouter()
	return &SynkrClient{ db: db, router: r }
}


