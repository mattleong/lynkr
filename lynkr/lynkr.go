package lynkr

import (
	"flag"
	"log"
	"net/http"
	"strings"
)

type LynkrClient struct {
	db *LynkrDB
	router *LynkrRouter
}

func NewLynkrClient() *LynkrClient {
	dbUri := getDBURI()
	db := newDBClient(dbUri)
	router := newRouter()
	client := LynkrClient{
		db: db,
		router: router,
	}
	return &client
}

func getPort() string {
	var port strings.Builder
	port.WriteString(*flag.String("port", "3000", "Localhost port (Default: 3000)"))
	return ":" + port.String()
}

func (s *LynkrClient) ServerStart() {
	log.Println("Lynkr server started...")

	s.router.setRoutes(s.db)
	httpErr := http.ListenAndServe(getPort(), nil)
	if httpErr != nil {
		return
	}
}

