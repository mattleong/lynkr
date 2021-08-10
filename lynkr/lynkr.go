package lynkr

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"strings"
)

type LynkrClient struct {
	db *LynkrDB
	router *LynkrRouter
}

func NewLynkrClient() *LynkrClient {
	db := newDBClient()
	r := newRouter()
	client := LynkrClient{ db: db, router: r }
	client.setRoutes()
	return &client
}

func getPort() string {
	var port strings.Builder
	port.WriteString(*flag.String("port", "3000", "Localhost port (Default: 3000)"))
	return ":" + port.String()
}

func (s *LynkrClient) ServerStart() {
	fmt.Println("Lynkr server started...")

	http.Handle("/", s.router.r)
	httpErr := http.ListenAndServe(getPort(), nil)
	if httpErr != nil {
		return
	}
}

func (s *LynkrClient) SaveLynk(ctx context.Context, requestLynk *RequestLynk) (*Lynk, error) {
	return s.db.saveLynk(ctx, requestLynk)
}

func (s *LynkrClient) FindLynkById(ctx context.Context, id string) (*Lynk, error) {
	return s.db.findLynkById(ctx, id)
}

