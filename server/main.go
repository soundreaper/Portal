package main

import (
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"github.com/soundreaper/portal/config"
	"github.com/soundreaper/portal/server"
)

func main() {
	// set up auth provider, the callback URL, and what user information to return
	goth.UseProviders(google.New(config.GetAuthConfig().GoogleKey, config.GetAuthConfig().GoogleSecret, "http://localhost:8000/auth/google/callback", "email"))

	// create a server with the default db connection
	s := server.NewServer(nil)

	// start the server on port :8000
	s.Start(":8000")
}
