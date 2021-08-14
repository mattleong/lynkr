package lynkr

import (
	"flag"
	"log"
	"net/http"
	"strings"
	"math/rand"
	"time"
	"github.com/mattleong/lynkr/pkg/routes"
	"github.com/mattleong/lynkr/pkg/db"
)

type LynkrClient struct {
	db db.DatabaseStore
	router *routes.Router
}

func NewLynkrClient() *LynkrClient {
	rand.Seed(time.Now().UnixNano())
	db := db.NewDBClient()
	router := routes.NewRouter()
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

	s.router.SetRoutes(s.db)
	httpErr := http.ListenAndServe(getPort(), nil)
	if httpErr != nil {
		return
	}
}

