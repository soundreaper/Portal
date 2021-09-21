package main

import (
	"github.com/soundreaper/portal/server"
)

func main() {
	// create a server with the default db connection
	s := server.NewServer(nil)

	// start the server on port :8080
	s.Start(":8080")
}
