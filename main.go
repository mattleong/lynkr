package main

import (
	"github.com/mattleong/lynkr/lynkr"
)

func main() {
	client := lynkr.NewLynkrClient()
	client.ServerStart()
}
