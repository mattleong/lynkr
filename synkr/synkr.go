package synkr

import (
	"context"
	"time"
)

type SynkrClient struct {
	db *SynkrDB
	router *SynkrRouter
}

func CreateContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	return ctx, cancel
}

func NewSynkrClient() *SynkrClient {
	db := NewDBClient()
	r := NewRouter()
	client := SynkrClient{ db: db, router: r }
	client.SetRoutes()
	return &client
}
