package lynkr

import (
	"fmt"
	"flag"
	"net/http"
	"strings"
)

func ServerStart() {
	fmt.Println("Lynkr server started...")

	// set up lynkr client
	s := NewLynkrClient()

	http.Handle("/", s.router.r)
	httpErr := http.ListenAndServe(getPort(), nil)
	if httpErr != nil {
		return
	}
}

func getPort() string {
	var port strings.Builder
	port.WriteString(*flag.String("port", "3000", "Localhost port (Default: 3000)"))
	return ":" + port.String()
}
