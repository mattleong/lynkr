package main

import (
	"github.com/mattleong/lynkr/pkg/client"
)

func main() {
	client := lynkr.NewLynkrClient()
	client.ServerStart()
}
