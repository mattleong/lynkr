package api

import (
	"fmt"
	"net/http"
	"github.com/mattleong/lynkr/logger"
)

func RootRoute(w http.ResponseWriter, req *http.Request) {
	logger.Log("hit: /")
	fmt.Fprintf(w, "hello\n")
}

func CreateRoute(w http.ResponseWriter, req *http.Request) {
	logger.Log("hit: /create")
	fmt.Fprintf(w, "create\n")
}
