package synkr

import (
	"context"
	"time"
)

type SynkrClient struct {
	db *SynkrDB
	router *SynkrRouter
}

func createContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	return ctx, cancel
}

func NewSynkrClient() *SynkrClient {
	db := newDBClient()
	r := newRouter()
	client := SynkrClient{ db: db, router: r }
	client.setRoutes()
	return &client
}
