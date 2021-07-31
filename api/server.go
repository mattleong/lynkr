package api

import (
	"fmt"
	"net/http"
)

func LynkrServerStart() {
	fmt.Println("init lynker")

	http.HandleFunc("/", RootRoute)
	http.HandleFunc("/create", RootRoute)
	http.ListenAndServe(":3000", nil)
}
