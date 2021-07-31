package api

import (
	"fmt"
	"net/http"
)

func RootRoute(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}
