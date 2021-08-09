package synkr

import (
	"fmt"
	"net/http"
)

func ServerStart() {
	fmt.Println("Lynkr server started...")

	// set up synkr client
	s := NewSynkrClient()

	http.Handle("/", s.router.r)
	httpErr := http.ListenAndServe(":3000", nil)
	if httpErr != nil {
		return
	}
}
