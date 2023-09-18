package main

import (
	"github.com/abrialstha/go-ewallet/api"
)

func main() {
	// Start the HttpServer
	port := 3000 // Later take this from say env
	server := api.NewServer(port)
	server.Start()
}
