package lynkr

import (
	"context"
	"time"
)

type LynkrClient struct {
	db *LynkrDB
	router *LynkrRouter
}

func createContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	return ctx, cancel
}

func NewLynkrClient() *LynkrClient {
	db := newDBClient()
	r := newRouter()
	client := LynkrClient{ db: db, router: r }
	client.setRoutes()
	return &client
}

func (s *LynkrClient) SaveLynk(ctx context.Context, requestLynk *RequestLynk) (*Lynk, error) {
	return s.db.saveLynk(ctx, requestLynk)
}

func (s *LynkrClient) FindLynkById(ctx context.Context, id string) (*Lynk, error) {
	return s.db.findLynkById(ctx, id)
}
