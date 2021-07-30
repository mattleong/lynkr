package main

import (
	"fmt"
	"net/http"
)

func rootRoute(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func main() {
	fmt.Println("init lynker")

	http.HandleFunc("/", rootRoute)
	http.ListenAndServe(":3000", nil)
}
