package main

import (
	"fmt"
	"net/http"
	"github.com/mattleong/lynkr/api"
)

func main() {
	fmt.Println("init lynker")

	http.HandleFunc("/", api.RootRoute)
	http.HandleFunc("/create", api.RootRoute)
	http.ListenAndServe(":3000", nil)
}
