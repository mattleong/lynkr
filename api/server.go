package api

import (
	"net/http"
	"github.com/mattleong/lynkr/logger"
)

func ServerStart() {
	logger.Log("Lynkr server started...")

	http.HandleFunc("/", RootRoute)
	http.HandleFunc("/create", CreateRoute)
	http.ListenAndServe(":3000", nil)
}
