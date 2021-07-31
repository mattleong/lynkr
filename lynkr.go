package main

import (
	"fmt"
	"net/http"
	"github.com/mattleong/lynkr/api/routes"
)

func main() {
	fmt.Println("init lynker")

	http.HandleFunc("/", RootRoute)
	http.HandleFunc("/create", RootRoute)
	http.ListenAndServe(":3000", nil)
}
