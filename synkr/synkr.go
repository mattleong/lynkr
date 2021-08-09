package synkr

import (
	"context"
	"time"
	"github.com/mattleong/lynkr/lynkr"
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

func (s *SynkrClient) SaveLynk(ctx context.Context, requestLynk *RequestLynk) (*lynkr.Lynk, error) {
	return s.db.saveLynk(ctx, requestLynk)
}

func (s *SynkrClient) FindLynkById(ctx context.Context, id string) (*lynkr.Lynk, error) {
	return s.db.findLynkById(ctx, id)
}
