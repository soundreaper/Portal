package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
	"github.com/soundreaper/portal/config"
)

// Jwks represents an array of json web keys
type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

// JSONWebKeys represents a json web key
type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

// getJwtMiddleware will return the jwt token middleware configuration
func getJwtMiddleware() jwtmiddleware.JWTMiddleware {
	return *jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			cert, err := getPemCert(token)
			if err != nil {
				log.Fatal(err)
			}

			result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
			return result, nil
		},
		SigningMethod: jwt.SigningMethodRS256,
	})
}

// getPemCert will get the pem certificate from the API
func getPemCert(token *jwt.Token) (string, error) {
	// Auth0 API information is pulled via the auth config function
	c := config.GetAuthConfig()
	cert := ""

	// Create url using the Auth0 ID pulled from config function
	url := fmt.Sprintf("%s/.well-known/jwks.json", c.Auth0ID)
	resp, err := http.Get(url)

	if err != nil {
		return cert, err
	}
	defer resp.Body.Close()

	var jwks = Jwks{}
	err = json.NewDecoder(resp.Body).Decode(&jwks)

	if err != nil {
		return cert, err
	}

	for k, _ := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		err := errors.New("Unable to find appropriate key.")
		return cert, err
	}

	return cert, nil
}
